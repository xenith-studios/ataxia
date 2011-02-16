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
    "ataxia/settings"
    //"io"
)


// Flag variables
var portFlag int
var configFlag string


/*
    Do some useful initialization here.
*/
func init() {
    // Setup the command-line flags
    flag.IntVar(&portFlag, "port", 0, "Main port")
    flag.StringVar(&configFlag, "config file", "data/config.lua", "Config file")

    // Parse the command line
    flag.Parse()

    // If previous shutdown was not clean, clean up state and environment
        // Are we recovering from a hotboot? Don't clean up network connections

    // Read configuration file

    // Initializations
        // Environment
        // Settings
        // Logging
        // Queues
        // Database
        // Lua
        // Network
}


// Perform a hotboot
// Save game and world state, save player state, save player list
func hotboot() {
    // Save game state
    // Save socket and player list
    os.Exec(os.Args[0], os.Args, os.Environ())
}


// Recover from a hotboot
// Restore game and world state, restore player list, restore player state
func restore() {
    
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
            fmt.Fprintln(os.Stderr, "Failed to chroot:", os.Errno(err))
            os.Exit(err)
        }        
    }
    
    // Write out pid file
    pid := fmt.Sprint(os.Getpid())
    pidfile, err := os.Open(settings.Pidfile, os.O_RDWR|os.O_CREAT, 0666)
    if pidfile == nil {
        fmt.Fprintln(os.Stderr, "Error writing pid to file:", err)
        os.Exit(1)
    }
    pidfile.Write([]byte(pid))
    pidfile.Close()
    defer os.Remove(settings.Pidfile)

    // Daemonize if configured

    // Initialize game state
        // Load database
        // Load commands
        // Load scripts
        // Load world
        // Load entities

    // Are we recovering from a hotboot?
        // Restore socket connections

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
