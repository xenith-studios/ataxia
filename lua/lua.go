/*
   Lua
*/
package lua

import (
	"log"
	golua "github.com/aarzilli/golua/lua"
	luar "github.com/stevedonovan/luar"
)

var MainState *golua.State

func testing(num string) {
	log.Printf("testing value %s", num)
}

func Initialize() bool {
	log.Println("Initializing Lua")

	if MainState != nil {
		return true
	}

	MainState = golua.NewState()
	MainState.OpenLibs()

	return true
}

// register gofuncs for interacting with game state
func Register(str string, f *func()) {
	luar.Register(MainState, "", luar.Map{
		str:f,
	})
}

func Shutdown() {
	if MainState != nil {
		MainState.Close()
	}
}

// simple command, one arg, no results
func Execute(func_name string, args string) {
	MainState.GetField(golua.LUA_GLOBALSINDEX, func_name)
	MainState.PushString(args)
	MainState.Call(1, 0)
} 

// two argument execute, passes executing player id
func ExecuteInterpret(func_name string, player_uid string, args string) {
	MainState.GetField(golua.LUA_GLOBALSINDEX, func_name)
	MainState.PushString(player_uid)
	MainState.PushString(args)
	MainState.Call(2, 0)
}

