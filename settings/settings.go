/*
   Settings
*/
package settings

import (
	"log"
	"github.com/xenith-studios/ataxia/lua"
)

var (
	Chroot    string = ""
	Pidfile   string = "data/ataxia.pid"
	Daemonize bool   = false
	MainPort  int    = 9000
)

// Pass the config file to Lua for processing, and pull the variables from Lua into package-level variables.
// Pass in each command-line flag for overriding values in the config file
func LoadConfigFile(configFile string, mainPort int) bool {
	L := lua.MainState
	err := L.DoFile(configFile)
	if err != nil {
		log.Fatal("Failed to read configuration file.")
		return false
	}

	L.GetGlobal("pid_file")
	Pidfile = L.CheckString(1)
	L.Pop(1)

	if mainPort != 0 {
		MainPort = mainPort
	} else {
		L.GetGlobal("main_port")
		MainPort = L.CheckInteger(1)
		L.Pop(1)
	}

	L.GetGlobal("chroot")
	Chroot = L.CheckString(1)
	L.Pop(1)

	L.GetGlobal("daemonize")
	Daemonize = L.ToBoolean(1)
	L.Pop(1)

	return true
}
