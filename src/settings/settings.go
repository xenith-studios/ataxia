/*
    Settings
*/
package settings

import (
    "log"
    "lua51"
    "ataxia/lua"
)

var (
    Chroot string = ""
    Pidfile string = "data/ataxia.pid"
    Daemonize bool = false
    MainPort int = 9000
)

func ParseConfigFile(configFile string, portFlag int) bool {
    L := lua.MainState
    ok := L.DoFile(configFile)
    if !ok {
        log.Fatal("Failed to read configuration file.")
        return false
    }

    lua.MainState.GetGlobal("pid_file")
    Pidfile = lua51.CheckString(L, 1)
    L.Pop(1)

    if portFlag != 0 {
        MainPort = portFlag
    } else {
        L.GetGlobal("main_port")
        MainPort = lua51.CheckInteger(L, 1)
        L.Pop(1)
    }
    
    L.GetGlobal("chroot")
    Chroot = lua51.CheckString(L, 1)
    L.Pop(1)

    L.GetGlobal("daemonize")
    Daemonize = L.ToBoolean(1)
    L.Pop(1)

    return true
}
