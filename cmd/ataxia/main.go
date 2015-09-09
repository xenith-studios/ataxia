/*
Ataxia Mud Engine

Copyright © 2009-2015 Xenith Studios
*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/naoina/toml"
	"github.com/xenith-studios/ataxia/engine"
)

// tomlConfig is the struct for parsing the TOML config file
type tomlConfig struct {
	MainPort          int
	AdminPort         int
	BuildPort         int
	PidFile           string
	LogFacility       string
	LogFile           string
	EmailFacility     string
	SendmailPath      string
	AdminEmail        string
	AbuseEmail        string
	StorageFacility   string
	DataPath          string
	MaxClients        int
	MaxClientsPerHost int
	AccountCreation   bool
	CharsPerAccount   int
	Autosave          int
	BackupCharacters  bool
}

// Variables for the command-line flags and config struct
var (
	configFlag     string
	hotbootFlag    bool
	descriptorFlag int
	mainPortFlag   int
	pidFlag        string
	config         tomlConfig
)

// When hotboot is called, this function will save game and world state, save each player state, and save the player list.
// Then it will do some cleanup (including closing the database) and call Exec to reload the running program.
// TODO: This is currently a stub function to lay out future functionality.
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
// TODO: This is currently a stub function to lay out future functionality.
func recover() {
	log.Println("Recovering from hotboot.")
}

func main() {
	fmt.Printf(`Ataxia Engine %s © 2009-2015, Xenith Studios (see AUTHORS)
Compiled on %s
Ataxia Engine comes with ABSOLUTELY NO WARRANTY; see LICENSE for details.
This is free software, and you are welcome to redistribute it
under certain conditions; for details, see the file LICENSE.

`, ataxiaVersion, ataxiaCompiled)

	// Setup the command-line flags (with defaults)
	flag.StringVar(&configFlag, "config", "data/config.toml", "Config file")
	flag.BoolVar(&hotbootFlag, "hotboot", false, "Recover from hotboot")
	flag.IntVar(&descriptorFlag, "descriptor", 0, "Hotboot descriptor")

	// Setup the flags that are defined in the config file but can be overriden
	// via the command-line
	flag.IntVar(&mainPortFlag, "main_port", 0, "Main engine port")
	flag.StringVar(&pidFlag, "pid_file", "data/ataxia.pid", "PID filename")

	// Parse the command line
	flag.Parse()

	// Parse the config file into the config struct and override values if specified on the command line
	f, err := os.Open(configFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	if err := toml.Unmarshal(buf, &config); err != nil {
		log.Fatal(err)
	}
	if mainPortFlag != 0 {
		config.MainPort = mainPortFlag
	}
	if pidFlag != "" {
		config.PidFile = pidFlag
	}
	log.Println("Loaded config file.")

	if !hotbootFlag {
		// If previous shutdown was not clean and we are not recovering from a hotboot, clean up state and environment if needed
	}

	// Initializations
	// Environment
	// Logging
	// Queues
	// Database

	// Create a channel that the engine can send a message on when it shuts down, so we can cleanup in the main goroutine
	shutdown := make(chan bool)

	// Spin up a goroutine to catch and handle signals
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

	// Write out pid file
	pid := fmt.Sprint(os.Getpid())
	pidfile, err := os.Create(config.PidFile)
	if err != nil {
		log.Fatalln("Error writing pid to file:", err)
	}
	pidfile.Write([]byte(pid))
	log.Println("Wrote PID to", config.PidFile)
	pidfile.Close()
	defer os.Remove(config.PidFile)

	// Initialize the game engine and network
	log.Println("Initializing the game engine")
	server := engine.NewServer(config.MainPort, shutdown)
	log.Println("Engine is running on port", config.MainPort)

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

	// Signal everything to cleanly shut down
	log.Println("Shutdown detected. Cleaning up....")
	server.Shutdown()
}
