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
    case "preview":
        startServer()
    case "build":
        build()
    case "upload":
        // TODO
    case "deploy":
        // TODO
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
    fmt.Println("build   - Build the website, overwriting the /public directory")
    fmt.Println("preview - Start a local server for previewing the website")
    fmt.Println("help    - Print this help text")
}

// Starts the server and launches the browser to view it
func startServer() {
    fmt.Println("=============================")
    fmt.Println("Starting local preview server")
    fmt.Println("=============================")

    cmd := exec.Command("hugo", "server")
    if err := cmd.Start(); err != nil {
        fmt.Println("Error starting command 'hugo server'")
        fmt.Println(err)
        return
    }

    if err := exec.Command("x-www-browser", "http://localhost:1313").Run(); err != nil {
        fmt.Println("Error running command 'x-www-browser http://localhost:1313'")
        fmt.Println(err)
        return
    }

    if err := cmd.Wait(); err != nil {
        fmt.Println("Error in command 'hugo server'")
        fmt.Println(err)
        return
    }
}

// Build the website, which places it in the /public directory
func build() {
    // Clear the /public directory to ensure clean build
    if err := os.RemoveAll("public"); err != nil {
       fmt.Println("Error clearing /public directory")
       return
    }

    // build the website
    if err := exec.Command("hugo").Run(); err != nil {
        fmt.Println("Error running command 'hugo'")
        fmt.Println(err)
        return
    }
}


//////
// Helpers
////////
