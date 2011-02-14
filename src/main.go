/*
    Ataxia Mud Engine

    Copyright © 2009-2010 Xenith Studios
*/
package main

import (
    "fmt"
    "flag"
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

    // Initializations
        // Settings
        // Logging
        // Queues
        // Database
        // Lua
        // Network
}


func main() {
    fmt.Println(`Ataxia Engine V0.1 © 2009-2011, Xenith Studios (see AUTHORS)
Ataxia Engine comes with ABSOLUTELY NO WARRANTY; see COPYING for details.
This is free software, and you are welcome to redistribute it
under certain conditions; for details, see the file COPYING.
`);

    // Parse the command line
    flag.Parse()

    // chroot

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
