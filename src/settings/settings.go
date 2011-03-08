package settings

import (
    "log"
    //"fmt"
    "lua51"
)

var Chroot string = ""
var Pidfile string = "data/ataxia.pid"
var Daemonize bool = false
var Port int = 9000

func ParseConfigFile(L *lua51.State, configFile string, portFlag int) bool {
    ok := L.DoFile(configFile)
    if !ok {
        log.Fatal("Failed to read configuration file.")
        return false
    }

    L.GetGlobal("pidfile")
    Pidfile = lua51.CheckString(L, 1)
    L.Pop(1)

    L.GetGlobal("port")
    Port = lua51.CheckInteger(L, 1)
    L.Pop(1)
    if portFlag != 0 {
        Port = portFlag
    }

    L.GetGlobal("chroot")
    Chroot = lua51.CheckString(L, 1)
    L.Pop(1)

    L.GetGlobal("daemonize")
    Daemonize = L.ToBoolean(1)
    L.Pop(1)

    return true
}
