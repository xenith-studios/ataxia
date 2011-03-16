/*
    Lua
*/
package lua

import (
    "log"
    "lua51"
)

var MainState *lua51.State

func Initialize() bool {
    log.Println("Initializing Lua")

    if MainState != nil {
        return true
    }

    MainState = lua51.NewState()
    MainState.OpenLibs()
    return true
}

func Shutdown() {
    if MainState != nil {
        MainState.Close()
    }
}