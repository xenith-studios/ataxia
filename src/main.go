/*
    Ataxia Mud Engine

    Copyright © 2009-2010 Xenith Studios
*/
package main

import (
    "fmt"
    "flag"
    "os"
    "syscall"
    "log"
    "lua51"
    "ataxia/settings"
)


// Flag variables
var portFlag int
var configFlag string
var hotbootFlag bool
var descriptorFlag int


// Main Lua State
var lua_state *lua51.State

/*
    Do some useful initialization here.
*/
func init() {
    // Setup the command-line flags
    flag.IntVar(&portFlag, "port", 0, "Main port")
    flag.StringVar(&configFlag, "config", "etc/config.lua", "Config file")
    flag.BoolVar(&hotbootFlag, "hotboot", false, "Recover from hotboot")
    flag.IntVar(&descriptorFlag, "descriptor", 0, "Hotboot descriptor")

    // Parse the command line
    flag.Parse()

    if !hotbootFlag {
        // If previous shutdown was not clean and we are not recovering from a hotboot, clean up state and environment
    }

    // Initialize Lua
    lua_state = lua51.NewState()
    lua_state.OpenLibs()

    // Read configuration file
    ok := settings.ParseConfigFile(lua_state, configFlag, portFlag)
    if !ok {
        log.Fatal("Error reading config file.")
    }

    // Initializations
        // Environment
        // Settings
        // Logging
        // Queues
        // Database
        // Network

    // Set up signal handlers
    fmt.Println(settings.Port)
}


// Perform a hotboot
// Save game and world state, save player state, save player list
func hotboot() {
    // Save game state
    // Save socket and player list
    // Disconnect from database
    arglist := append(os.Args, "-hotboot", "-descriptor=", fmt.Sprint(1234))
    os.Exec(os.Args[0], arglist, os.Environ())

    // If we get to this point, something went wrong. Die.
    log.Fatal("Failed to exec during hotboot.")
}


// Recover from a hotboot
// Restore game and world state, restore player list, restore player state
func recover() {
    fmt.Print("Recovering from hotboot.")
}


func main() {
    fmt.Println(`Ataxia Engine V0.1 © 2009-2011, Xenith Studios (see AUTHORS)
Ataxia Engine comes with ABSOLUTELY NO WARRANTY; see COPYING for details.
This is free software, and you are welcome to redistribute it
under certain conditions; for details, see the file COPYING.
`);

    // chroot into the configured directory
    if settings.Chroot != "" {
        err := syscall.Chroot(settings.Chroot)
        if err != 0 {
            log.Fatalln("Failed to chroot:", os.Errno(err))
        }
    }

    // Write out pid file
    pid := fmt.Sprint(os.Getpid())
    pidfile, err := os.Open(settings.Pidfile, os.O_RDWR|os.O_CREAT, 0666)
    if pidfile == nil {
        log.Fatalln("Error writing pid to file:", err)
    }
    pidfile.Write([]byte(pid))
    pidfile.Close()
    defer os.Remove(settings.Pidfile)

    // Daemonize if configured
    if settings.Daemonize {
        log.Println("Daemonizing")
        // Daemonize here
    }

    // Initialize game state
        // Load database
        // Load commands
        // Load scripts
        // Load world
        // Load entities

    // Are we recovering from a hotboot?
        // Restore socket connections

    // Run the game loop in its own goroutine
    // Main loop
        // Handle network messages (push user events)
        // Handle game updates
            // Game tick
            // Time update
            // Weather update
            // Entity updates (push events)
        // Handle pending events
        // Handle pending messages (network and player)
        // Sleep

   // Cleanup
}
