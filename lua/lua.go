/*
   Lua
*/
package lua

import (
	"log"
	golua "github.com/aarzilli/golua/lua"
)

var MainState *golua.State

func Initialize() bool {
	log.Println("Initializing Lua")

	if MainState != nil {
		return true
	}

	MainState = golua.NewState()
	MainState.OpenLibs()
	return true
}

func Shutdown() {
	if MainState != nil {
		MainState.Close()
	}
}

// simple command, one arg, no results
func Execute(func_name string, args string) {
	// need to lock the vm here
	MainState.GetField(golua.LUA_GLOBALSINDEX, func_name)
	MainState.PushString(args)
	MainState.Call(1, 0)
} 
