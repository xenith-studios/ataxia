/*
    Settings
*/
package settings

import (
    "log"
    "github.com/xenith-studios/golua"
    "ataxia/lua"
)

var (
    Chroot string = ""
    Pidfile string = "data/ataxia.pid"
    Daemonize bool = false
    MainPort int = 9000
)

// Pass the config file to Lua for processing, and pull the variables from Lua into package-level variables.
// Pass in each command-line flag for overriding values in the config file
func LoadConfigFile(configFile string, mainPort int) bool {
    L := lua.MainState
    ok := L.DoFile(configFile)
    if !ok {
        log.Fatal("Failed to read configuration file.")
        return false
    }

    lua.MainState.GetGlobal("pid_file")
    Pidfile = golua.CheckString(L, 1)
    L.Pop(1)

    if mainPort != 0 {
        MainPort = mainPort
    } else {
        L.GetGlobal("main_port")
        MainPort = golua.CheckInteger(L, 1)
        L.Pop(1)
    }
    
    L.GetGlobal("chroot")
    Chroot = golua.CheckString(L, 1)
    L.Pop(1)

    L.GetGlobal("daemonize")
    Daemonize = L.ToBoolean(1)
    L.Pop(1)

    return true
}
