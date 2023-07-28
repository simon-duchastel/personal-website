package main

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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
	var err error
	switch command := args[0]; command {
	case "preview": // preview the website on a local server
		err = startServer()
	case "build": // build the website
		err = build()
	case "upload": // upload the website
		err = upload()
	case "deploy": // convenience command for build + upload
		if err = build(); err == nil {
			err = upload() // only upload if there was no error building
		}
	case "rotatecert": // rotate the ssl (https) cert for the website
		err = rotateCert()
	case "help":
		printHelp()
		err = errors.New("No command run")
	default:
		fmt.Println("Error: Invalid command '" + command + "'")
		fmt.Println()
		printHelp()
		err = errors.New("Invalid command")
	}
	if err == nil {
		fmt.Println("Success!")
	} else {
		fmt.Println(err)
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
	fmt.Println("  build       - Build the website, overwriting the " + HUGO_BUILD_DIRECTORY + " directory")
	fmt.Println("  deploy      - Build and upload the latest version of the website to simon.duchastel.com")
	fmt.Println("  preview     - Start a local server for previewing the website")
	fmt.Println("  upload      - Upload the built website and host it at simon.duchastel.com")
	fmt.Println("  rotatecert  - Rotate the ssl (https) cert for simon.duchastel.com, duchastel.com, and duchastel.org")
	fmt.Println("  help        - Print this help text")
}

// Starts the server and launches the browser to view it
func startServer() error {
	fmt.Println("Starting local preview server...")

	cmd := exec.Command("hugo", "server")
	if err := cmd.Start(); err != nil {
		fmt.Println("Error: cannot run command 'hugo server'")
		return err
	}

	if err := exec.Command("x-www-browser", "http://localhost:1313").Run(); err != nil {
		fmt.Println("Error: cannot run command 'x-www-browser http://localhost:1313'")
		return err
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error: cannot run command 'hugo server'")
		return err
	}

	return nil
}

// Build the website, which places it in the public/ directory
func build() error {
	// clear the public/ directory to ensure clean build
	fmt.Println("Clearing " + HUGO_BUILD_DIRECTORY + " directory")
	if err := os.RemoveAll(HUGO_BUILD_DIRECTORY); err != nil {
		fmt.Println("Error: cannot clear " + HUGO_BUILD_DIRECTORY + " directory")
		return err
	}

	// build the website
	fmt.Println("Building website")
	if err := exec.Command("hugo").Run(); err != nil {
		fmt.Println("Error: cannot run command 'hugo'")
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
		return err
	}
	defer client.Close()

	// root of the website, ie. /home/[username]/public_html/simon.duchastel.com
	websiteRoot := "/home/" + config.clientConfig.User + "/public_html/simon.duchastel.com"

	// copy old website to bin/site-old/ as a back-up in case there's some kind of issue
	fmt.Println("Copying old website from webhost to bin/website-old/ in case there are any issues")
	files, err := listRemoteFiles(client, websiteRoot)
	if err != nil {
		fmt.Println("Error: could not recursively list files from '" + websiteRoot + "'")
		return err
	}

	// clear bin/website-old/ directory in preparation for storing old website
	// if err := os.RemoveAll(SITE_OLD_DIRECTORY); err != nil {
	// 	fmt.Println("Error: cannot clear '" + SITE_OLD_DIRECTORY + "' directory")
	// 	return err
	// }

	if len(files) <= 0 {
		fmt.Println("    Nothing to download")
	}
	for _, file := range files {
		destinationFile := SITE_OLD_DIRECTORY + "/" + strings.TrimPrefix(file, websiteRoot+"/")

		fmt.Println("Downloading " + file + " to " + destinationFile)
		if err := downloadRemoteFile(client, file, destinationFile); err != nil {
			return err
		}
	}

	fmt.Println("Removing old website from webhost")
	// runRemoteCommandToConsole(client, "rm -rf "+websiteRoot+"/*") // delete everything in the website directory

	fmt.Println("Uploading new website to webhost")
	if err := filepath.WalkDir(HUGO_BUILD_DIRECTORY, func(path string, file fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			fmt.Println("Error: failed to read '" + path + "'")
			return err
		}

		if !file.IsDir() && len(path) > 0 {
			uploadFilePath := websiteRoot + "/" + strings.TrimPrefix(path, HUGO_BUILD_DIRECTORY+"/")
			fmt.Println("Uploading " + path + " to " + uploadFilePath)
			uploadFile(client, path, uploadFilePath)
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

// Rotate the ssl (https) cert for the simon.duchastel.com, duchastel.com, and
// duchastel.org domains
func rotateCert() error {
	fmt.Println("Command not yet implemented. Sorry!")

	return errors.New("Not yet implemented")
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

// Run a command on the remote host via ssh and return its output as a
// byte buffer
func runRemoteCommand(client *ssh.Client, command string) (*bytes.Buffer, error) {
	// start an interactive session
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("Error: failed to create session")
		return nil, err
	}
	defer session.Close()

	// execute a command on the session
	var buffer bytes.Buffer
	session.Stdout = &buffer
	if err := session.Run(command); err != nil {
		fmt.Println("Error: failed to run command '" + command + "'")
		return nil, err
	}

	return &buffer, nil
}

// Run a command on the remote host via ssh and print its output to console
func runRemoteCommandToConsole(client *ssh.Client, command string) error {
	buffer, err := runRemoteCommand(client, command)
	if err != nil {
		return err
	}
	fmt.Println(buffer.String())

	return nil
}

// Returns true if the remote file is confirmed to
// be a file, false otherwise
// Implemented by running the `test -f` command remotely
func remoteFileIsFile(client *ssh.Client, filePath string) (bool, error) {
	buffer, err := runRemoteCommand(client, "test -f "+filePath+" && echo true || echo false")
	if err != nil {
		return false, err
	}

	return strings.TrimSpace(buffer.String()) == "true", nil
}

// Returns true if the remote file is confirmed to
// be a directory, false otherwise.
// Implemented by running the `test -d` command remotely.
func remoteFileIsDirectory(client *ssh.Client, filePath string) (bool, error) {
	buffer, err := runRemoteCommand(client, "test -d "+filePath+" && echo true || echo false")
	if err != nil {
		return false, err
	}

	return strings.TrimSpace(buffer.String()) == "true", nil
}

// Upload a file with the given ssh client.
// [filePath] is the path to the file (including filename)
// [destinationFilePath] is the path to the file on the remote host (including filename)
func uploadFile(client *ssh.Client, filePath, destinationFilePath string) error {
	scpClient, err := scp.NewClientBySSH(client)
	if err != nil {
		fmt.Println("Error: failed to create file-transfer client")
		return err
	}

	if err := scpClient.Connect(); err != nil {
		fmt.Println("Error: failed to create file-transfer connection over ssh")
		return err
	}
	defer scpClient.Close()

	fileToUpload, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error: unable to open file '" + filePath + "'")
		return err
	}
	defer fileToUpload.Close()

	if err := scpClient.CopyFromFile(context.Background(), *fileToUpload, destinationFilePath, READ_ONLY_FILE); err != nil {
		fmt.Println("Error: failed to copy file ' " + filePath + "' to remote server")
		return err
	}

	return nil
}

// Download a file with the given ssh client. Downloads the file from
// [remoteFileLocation] and saves it as [destinationFileLocation] locally
// (destination location must include filename).
func downloadRemoteFile(client *ssh.Client, remoteFileLocation, destinationFileName string) error {
	scpClient, err := scp.NewClientBySSH(client)
	if err != nil {
		fmt.Println("Error: failed to create file-transfer client")
		return err
	}

	if err := scpClient.Connect(); err != nil {
		fmt.Println("Error: failed to create file-transfer connection over ssh")
		return err
	}
	defer scpClient.Close()

	file, err := createFileWithDirectories(destinationFileName)
	defer file.Close()

	scpClient.CopyFromRemote(context.Background(), file, remoteFileLocation)

	return nil
}

// List all files, including hidden ones (but not directories) within
// the remote directory specified by [remoteDirectoryPath]. Includes
// recursive files, ie. listing remote files in /foo will list /foo/bar/baz.txt
func listRemoteFiles(client *ssh.Client, remoteDirectoryPath string) ([]string, error) {
	buffer, err := runRemoteCommand(client, "ls -A "+remoteDirectoryPath)
	if err != nil {
		return nil, err
	}
	if buffer == nil || buffer.Len() <= 0 {
		return nil, nil // if buffer is nil or empty, no files were found
	}

	// each file is on its own line
	// ignore empty strings and recurse on directories
	files := strings.Split(buffer.String(), "\n")
	var filesToReturn []string
	for _, file := range files {
		// clean up the file name and skip any blank files
		trimmedFile := strings.TrimSpace(file)
		if len(trimmedFile) <= 0 {
			continue
		}
		fullFileName := remoteDirectoryPath + "/" + trimmedFile

		isFile, err := remoteFileIsFile(client, fullFileName)
		if err != nil {
			return nil, err
		}

		isDirectory, err := remoteFileIsDirectory(client, fullFileName)
		if err != nil {
			return nil, err
		}

		if isFile {
			filesToReturn = append(filesToReturn, fullFileName)
		}
		if isDirectory {
			recursiveFiles, err := listRemoteFiles(client, fullFileName)
			if err != nil {
				return nil, err
			}
			filesToReturn = append(filesToReturn, recursiveFiles...)
		}
	}

	return filesToReturn, nil
}

// Create the file if it does not exist, as well as all
// intermediate directories if they do not exist
func createFileWithDirectories(filePath string) (*os.File, error) {
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		fmt.Println("Error: unable to create directories for '" + filePath + "'")
		return nil, err
	}
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error: unable to create file '" + filePath + "'")
		return nil, err
	}

	return file, nil
}

////////
// Constants
////////

// Flag to use in config to instruct insecure ssh connection
const INSECURE_MODE = "insecure"

// Flag to use for setting file as read-only on the file system
const READ_ONLY_FILE = "0644"

// Location to store old website as backup while uploading/deploying new website
const SITE_OLD_DIRECTORY = "bin/website-old"

// Location of the Hugo build output directory
const HUGO_BUILD_DIRECTORY = "public"
