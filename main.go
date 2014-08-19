/*
   Ataxia Mud Engine

   Copyright © 2009-2014 Xenith Studios
*/
package main

import (
	"flag"
	"fmt"
	"github.com/xenith-studios/ataxia/engine"
	"github.com/xenith-studios/ataxia/lua"
	"github.com/xenith-studios/ataxia/settings"
	"log"
	"os"
	"os/signal"
	"syscall"
	//	log "log4go.googlecode.com/hg"
)

// Variables for the command-line flags
var (
	portFlag       int
	configFlag     string
	hotbootFlag    bool
	descriptorFlag int
)

var shutdown chan bool

// Do all our basic initialization within the main package's init function.
func init() {
	fmt.Printf(`Ataxia Engine %s © 2009-2014, Xenith Studios (see AUTHORS)
Compiled on %s
Ataxia Engine comes with ABSOLUTELY NO WARRANTY; see COPYING for details.
This is free software, and you are welcome to redistribute it
under certain conditions; for details, see the file COPYING.

`, ATAXIA_VERSION, ATAXIA_COMPILED)

	shutdown = make(chan bool)
	// Setup the command-line flags
	flag.IntVar(&portFlag, "port", 0, "Main port")
	flag.StringVar(&configFlag, "config", "etc/config.lua", "Config file")
	flag.BoolVar(&hotbootFlag, "hotboot", false, "Recover from hotboot")
	flag.IntVar(&descriptorFlag, "descriptor", 0, "Hotboot descriptor")

	// Parse the command line
	flag.Parse()

	// Initialize Lua
	lua.MainState = lua.NewState()

	// Read configuration file
	ok := settings.LoadConfigFile(configFlag, portFlag)
	if !ok {
		log.Fatal("Error reading config file.")
	}
	log.Println("Loaded config file.")

	// Initializations
	// Environment
	// Logging
	// Queues
	// Database

	if !hotbootFlag {
		// If previous shutdown was not clean and we are not recovering from a hotboot, clean up state and environment if needed
	}
}

// When hotboot is called, this function will save game and world state, save each player state, and save the player list.
// Then it will do some cleanup (including closing the database) and call Exec to reload the running program.
func hotboot() {
	// Save game state
	// Save socket and player list
	// Disconnect from database
	arglist := append(os.Args, "-hotboot", "-descriptor=", fmt.Sprint(1234))
	syscall.Exec(os.Args[0], arglist, os.Environ())

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
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	go func() {
		for sig := range c {
			if usig, ok := sig.(os.Signal); ok {
				switch usig {
				case syscall.SIGQUIT:
					fallthrough
				case syscall.SIGTERM:
					fallthrough
				case syscall.SIGINT:
					// Catch the three interrupt signals and signal the game to shutdown.
					log.Println("Received SIGINT, shutting down.")
					shutdown <- true
				case syscall.SIGHUP:
					// TODO: Reload settings and game state
					log.Println("Received SIGHUP, reloading configuration and game state.")
				}
			}
		}
	}()

	// If configured, chroot into the designated directory
	if settings.Chroot != "" {
		err := syscall.Chroot(settings.Chroot)
		if err != nil {
			log.Fatalln("Failed to chroot:", err)
		}
		err = os.Chdir(settings.Chroot)
		if err != nil {
			log.Fatalln("Failed to chdir:", err)
		}
		log.Println("Chrooted to", settings.Chroot)
	}

	// Drop priviledges if configured

	// Daemonize if configured
	if settings.Daemonize {
		log.Println("Daemonize functionality is currently disabled. Continuing as normal...")
		// log.Println("Daemonizing...")
		// Daemonize here
		// TODO: This probably won't be doable until Go supports forking into the background.
	}

	// Write out pid file
	pid := fmt.Sprint(os.Getpid())
	pidfile, err := os.Create(settings.Pidfile)
	if err != nil {
		log.Fatalln("Error writing pid to file:", err)
	}
	pidfile.Write([]byte(pid))
	log.Println("Wrote PID to", settings.Pidfile)
	pidfile.Close()
	defer os.Remove(settings.Pidfile)

	// Initialize the network
	log.Println("Initializing network")
	server := engine.NewServer(settings.MainPort, shutdown)
	log.Println("Server running on port", settings.MainPort)

	// at this point, server and world go functions have been published
	// to Lua, we can load up some libraries for scripting action
	err = lua.MainState.DoFile("scripts/interface/context.lua")
	if err != nil {
		log.Fatal(err)
	}
	err = lua.MainState.DoFile("scripts/interface/accessors.lua")
	if err != nil {
		log.Fatal(err)
	}
	err = lua.MainState.DoFile("scripts/interface/character.lua")
	if err != nil {
		log.Fatal(err)
	}
	err = lua.MainState.DoFile("scripts/interface/room.lua")
	if err != nil {
		log.Fatal(err)
	}
	err = lua.MainState.DoFile("scripts/commands/character_action.lua")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize game state
	// Load database

	// Load commands

	// Load scripts
	// Load world

	// Load entities

	server.InitializeWorld()

	// Are we recovering from a hotboot?
	if hotbootFlag {
		recover()
	}

	// Initialization and setup is complete. Spin up a goroutine to handle incoming connections
	go server.Listen()

	// Run the game loop in its own goroutine
	go server.Run()

	// Wait for the shutdown signal
	<-shutdown

	// Cleanup
	log.Println("Shutdown detected. Cleaning up....")
	lua.Shutdown(lua.MainState)
	server.Shutdown()
}
