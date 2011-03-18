/*
    Ataxia Mud Engine

    Copyright © 2009-2010 Xenith Studios
*/
package main

import (
    "fmt"
    "flag"
    "os"
    "os/signal"
    "syscall"
    "log"
    "ataxia/lua"
    "ataxia/settings"
)


// Variables for the command-line flags
var (
    portFlag int
    configFlag string
    hotbootFlag bool
    descriptorFlag int
)

var shutdown chan bool

// Do all our basic initialization within the main package's init function.
func init() {
    fmt.Println(`Ataxia Engine V0.1 © 2009-2011, Xenith Studios (see AUTHORS)
Ataxia Engine comes with ABSOLUTELY NO WARRANTY; see COPYING for details.
This is free software, and you are welcome to redistribute it
under certain conditions; for details, see the file COPYING.
`);

    shutdown = make(chan bool)
    // Setup the command-line flags
    flag.IntVar(&portFlag, "port", 0, "Main port")
    flag.StringVar(&configFlag, "config", "etc/config.lua", "Config file")
    flag.BoolVar(&hotbootFlag, "hotboot", false, "Recover from hotboot")
    flag.IntVar(&descriptorFlag, "descriptor", 0, "Hotboot descriptor")

    // Parse the command line
    flag.Parse()

    // Initialize Lua
    lua.Initialize()
 
    // Read configuration file
    ok := settings.LoadConfigFile(configFlag, portFlag)
    if !ok {
        log.Fatal("Error reading config file.")
    }

    // Initializations
        // Environment
        // Logging
        // Queues
        // Database

    // Initialize the network
    log.Println("Initializing network")
    InitializeNetwork()

    if !hotbootFlag {
        // If previous shutdown was not clean and we are not recovering from a hotboot, clean up state and environment
    }
}


// When hotboot is called, this function will save game and world state, save each player state, and save the player list.
// Then it will do some cleanup (including closing the database) and call Exec to reload the running program.
func hotboot() {
    // Save game state
    // Save socket and player list
    // Disconnect from database
    arglist := append(os.Args, "-hotboot", "-descriptor=", fmt.Sprint(1234))
    os.Exec(os.Args[0], arglist, os.Environ())

    // If we get to this point, something went wrong. Die.
    log.Fatal("Failed to exec during hotboot.")
}


// When recovering from a hotboot, recover will restore the game and world state, restore the player list, and restore each player state.
// Once that is done, it will then reconnect each active descriptor to the associated player.
func recover() {
    log.Println("Recovering from hotboot.")
}


// 
func main() {
    // At this point, basic initialization has completed

    // Spin up a goroutine to handle signals
    go func() {
        for {
            select { 
                case sig := <- signal.Incoming:
                    switch sig.(signal.UnixSignal) {
                        case syscall.SIGTERM, syscall.SIGINT:
                            shutdown <- true
                    }
            }
        }
    }()

    // If configured, chroot into the designated directory
    if settings.Chroot != "" {
        err := syscall.Chroot(settings.Chroot)
        if err != 0 {
            log.Fatalln("Failed to chroot:", os.Errno(err))
        }
        error := os.Chdir("/")
        if error != nil {
            log.Fatalln("Failed to chdir:", error)
        }
        log.Println("Chrooted to", settings.Chroot)
    }

    // Write out pid file
    pid := fmt.Sprint(os.Getpid())
    pidfile, err := os.Open(settings.Pidfile, os.O_RDWR|os.O_CREAT, 0666)
    if pidfile == nil {
        log.Fatalln("Error writing pid to file:", err)
    }
    pidfile.Write([]byte(pid))
    log.Println("Wrote PID to", settings.Pidfile)
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
    if hotbootFlag {
        recover()
    }
    
    // Initialization and setup is complete. Spin up a goroutine to handle incoming connections
    go mainServer.Listen()

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
    
    // Wait for the shutdown signal
    <-shutdown

    // Cleanup
    log.Println("Cleaning up....")
    lua.Shutdown()
    mainServer.Shutdown()
}
