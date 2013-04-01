/*
   Lua
*/
package lua

import (
	"log"
	"github.com/xenith-studios/golua"
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
