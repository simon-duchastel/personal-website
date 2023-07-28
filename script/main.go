package main

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"

	scp "github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func main() {
	args := os.Args[1:] // ignore the first arg (program name)
	if len(args) != 1 {
		fmt.Println("Error: Must provide exactly one command")
		fmt.Println()
		printHelp()
		return
	}

	// parse the args and execute relevant command
	switch command := args[0]; command {
	case "preview": // preview the website on a local server
		startServer()
	case "build": // build the website
		build()
	case "upload": // upload the website
		upload()
	case "deploy": // convenience command for build + upload
		if build() == nil {
			upload() // only upload if there was no error building
		}
	case "rotatecert": // rotate the ssl (https) cert for the website
		rotateCert()
	case "help":
		printHelp()
	default:
		fmt.Println("Error: Invalid command '" + command + "'")
		fmt.Println()
		printHelp()
	}
}

//////
// Commands
////////

func printHelp() {
	programName := os.Args[0] // first arg is program name
	fmt.Println("Usage: " + programName + " [command]")

	fmt.Println()

	fmt.Println("Commands:")
	fmt.Println("  build       - Build the website, overwriting the /public directory")
	fmt.Println("  deploy      - Build and upload the latest version of the website to simon.duchastel.com")
	fmt.Println("  preview     - Start a local server for previewing the website")
	fmt.Println("  upload      - Upload the built website and host it at simon.duchastel.com")
	fmt.Println("  rotatecert  - Rotate the ssl (https) cert for simon.duchastel.com, duchastel.com, and duchastel.org")
	fmt.Println("  help    - Print this help text")
}

// Starts the server and launches the browser to view it
func startServer() error {
	fmt.Println("Starting local preview server...")

	cmd := exec.Command("hugo", "server")
	if err := cmd.Start(); err != nil {
		fmt.Println("Error: cannot run command 'hugo server'")
		fmt.Println(err)
		return err
	}

	if err := exec.Command("x-www-browser", "http://localhost:1313").Run(); err != nil {
		fmt.Println("Error: cannot run command 'x-www-browser http://localhost:1313'")
		fmt.Println(err)
		return err
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error: cannot run command 'hugo server'")
		fmt.Println(err)
		return err
	}

	return nil
}

// Build the website, which places it in the /public directory
func build() error {
	// clear the /public directory to ensure clean build
	fmt.Println("Clearing /public directory")
	if err := os.RemoveAll("public"); err != nil {
		fmt.Println("Error: cannot clear /public directory")
		return err
	}

	// build the website
	fmt.Println("Building website")
	if err := exec.Command("hugo").Run(); err != nil {
		fmt.Println("Error: cannot run command 'hugo'")
		fmt.Println(err)
		return err
	}

	return nil
}

// Upload the website to the webhost
func upload() error {
	fmt.Println("Connecting to webhost")
	// get the login configuration
	config, err := getSshClientConfig()
	if err != nil {
		return err
	}

	// start the connection to the webhost
	client, err := ssh.Dial("tcp", config.tcpAddress, config.clientConfig)
	if err != nil {
		fmt.Println("Error: failed to connect to webhost")
		fmt.Println(err)
		return err
	}
	defer client.Close()

	runRemoteCommand(client, "ls")

	// TODO - clear out public_html/simon.duchastel.com directory
	fmt.Println("Removing old website on webhost")

	// TODO - upload hugo site to public_html/simon.duchastel.com directory

	return nil
}

// Rotate the ssl (https) cert for the simon.duchastel.com, duchastel.com, and
// duchastel.org domains
func rotateCert() error {
	if err := exec.Command("hugo").Run(); err != nil {
		fmt.Println("Error: cannot run command 'hugo'")
		fmt.Println(err)
		return err
	}

	return nil
}

//////
// Helpers
////////

// Helper struct to hold all ssh config information
type sshConfig struct {
	clientConfig *ssh.ClientConfig
	tcpAddress   string
}

// Get ssh config from local ssh.config file
// ssh.config file MUST NOT be source-controlled (contains
// sensitive info like username/password)
func getSshClientConfig() (*sshConfig, error) {
	configFile, err := os.Open("ssh.config")
	defer configFile.Close()

	if err != nil {
		fmt.Println("Error: ssh config (username, password) must be provided in file ssh.config")
		fmt.Println("ssh.config format:")
		fmt.Println("- 1st line: username to auth into webhost ssh")
		fmt.Println("- 2nd line: password to auth into webhost ssh")
		fmt.Println("- 3rd line: tcp address in the format '[address]:[port]' (ex. 'server.com:22')")
		fmt.Println("- 4th line: location of ssh known_hosts file OR 'insecure' if host key should not be validated (INSECURE)")
		fmt.Println(err)
		return nil, err
	}

	fileScanner := bufio.NewScanner(configFile)
	fileScanner.Split(bufio.ScanLines)

	if !fileScanner.Scan() {
		fmt.Println("Error: 1st line of ssh.config must contain ssh username")
		return nil, errors.New("ssh config error")
	}
	username := fileScanner.Text()

	if !fileScanner.Scan() {
		fmt.Println("Error: 2nd line of ssh.config must contain ssh password")
		return nil, errors.New("ssh config error")
	}
	pasword := fileScanner.Text()

	if !fileScanner.Scan() {
		fmt.Println("Error: 3rd line of ssh.config must contain tcp address in the format '[address]:[port]' (ex: 'server.com:22')")
		return nil, errors.New("ssh config error")
	}
	tcpAddress := fileScanner.Text()

	if !fileScanner.Scan() {
		fmt.Println("Error: 4th line of ssh.config must either be file location of ssh known_hosts file OR '" +
			INSECURE_MODE + "' if INSECURE mode should be used (no host key validation)")

		return nil, errors.New("ssh config error")
	}

	var hostKeyCallback ssh.HostKeyCallback
	knownHosts := fileScanner.Text()
	if knownHosts == INSECURE_MODE {
		hostKeyCallback = ssh.InsecureIgnoreHostKey()
	} else {
		var err error
		hostKeyCallback, err = knownhosts.New(knownHosts)
		if err != nil {
			fmt.Println("Error: problem parsing ssh known_hosts file")
			fmt.Println(err)
			return nil, err
		}
	}

	return &sshConfig{
		&ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{
				ssh.Password(pasword),
			},
			HostKeyCallback: hostKeyCallback,
		},
		tcpAddress,
	}, nil
}

// Run a command on the remote host via ssh and print its output to console
func runRemoteCommand(client *ssh.Client, command string) error {
	// start an interactive session
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("Error: failed to create session")
		fmt.Println(err)
		return err
	}
	defer session.Close()

	// execute a command on the session
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(command); err != nil {
		fmt.Println("Error: failed to run command '" + command + "'")
		fmt.Println(err)
		return err
	}
	fmt.Println(b.String())

	return nil
}

// Upload a file with the given ssh client.
// [pathToFile] is the path to the file, [destinationFilePath] is the path
// to the file on the remote host (including filename)
func uploadFile(client *ssh.Client, pathToFile, destinationFilePath string) error {
	scpClient, err := scp.NewClientBySSH(client)
	if err != nil {
		fmt.Println("Error: failed to create file-transfer client")
		fmt.Println(err)
		return err
	}

	if err := scpClient.Connect(); err != nil {
		fmt.Println("Error: failed to create file-transfer connection over ssh")
		fmt.Println(err)
		return err
	}

	fileToUpload, _ := os.Open(pathToFile)
	defer fileToUpload.Close()
	defer scpClient.Close()

	if err := scpClient.CopyFromFile(context.Background(), *fileToUpload, destinationFilePath, READ_ONLY_FILE); err != nil {
		fmt.Println("Error: failed to copy file to remote server")
		fmt.Println(err)
		return err
	}

	return nil
}

////////
// Constants
////////

const INSECURE_MODE = "insecure"
const READ_ONLY_FILE = "0644"
