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
    fmt.Println("Usage: website [command]")

    fmt.Println()
    
    fmt.Println("Commands:")
    fmt.Println("help    - Print this help text")
    fmt.Println("preview - Start a local server for previewing the website")
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
    }

    if err := exec.Command("x-www-browser", "http://localhost:1313").Run(); err != nil {
        fmt.Println("Error running command 'x-www-browser https://localhost:1313'")
        fmt.Println(err)
    }

    if err := cmd.Wait(); err != nil {
        fmt.Println("Error in command 'hugo server'")
        fmt.Println(err)
    }
}


//////
// Helpers
////////
