package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    args := os.Args[1:] // ignore the first arg (program name)
    if len(args) != 1 {
        fmt.Println("Must provide exactly one command")
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
        if build() != nil {
            upload() // only upload if there was no error building
        }
    case "help":
        printHelp()
    default:
        fmt.Println("Invalid command '" + command + "'")
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
    fmt.Println("  build   - Build the website, overwriting the /public directory")
    fmt.Println("  deploy  - Build and upload the latest version of the website to simon.duchastel.com")
    fmt.Println("  preview - Start a local server for previewing the website")
    fmt.Println("  upload  - Upload the built website and host it at simon.duchastel.com")
    fmt.Println("  help    - Print this help text")
}

// Starts the server and launches the browser to view it
func startServer() error {
    fmt.Println("=============================")
    fmt.Println("Starting local preview server")
    fmt.Println("=============================")

    cmd := exec.Command("hugo", "server")
    if err := cmd.Start(); err != nil {
        fmt.Println("Error starting command 'hugo server'")
        fmt.Println(err)
        return err
    }

    if err := exec.Command("x-www-browser", "http://localhost:1313").Run(); err != nil {
        fmt.Println("Error running command 'x-www-browser http://localhost:1313'")
        fmt.Println(err)
        return err
    }

    if err := cmd.Wait(); err != nil {
        fmt.Println("Error in command 'hugo server'")
        fmt.Println(err)
        return err
    }

    return nil
}

// Build the website, which places it in the /public directory
func build() error {
    // Clear the /public directory to ensure clean build
    if err := os.RemoveAll("public"); err != nil {
       fmt.Println("Error clearing /public directory")
       return err
    }

    // build the website
    if err := exec.Command("hugo").Run(); err != nil {
        fmt.Println("Error running command 'hugo'")
        fmt.Println(err)
        return err
    }

    return nil
}

// Upload the website to the webserver
func upload() error {
    return nil
}


//////
// Helpers
////////
