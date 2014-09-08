package lua

import (
	"log"

	golua "github.com/aarzilli/golua/lua"
	//	luar "github.com/stevedonovan/luar"
)

// MainState is the main LuaState for the engine
var MainState *golua.State

// NewState returns a newly initalized LuaState
func NewState() *golua.State {
	log.Println("Initializing Lua State")
	st := golua.NewState()
	st.OpenLibs()
	return st
}

// Shutdown closes the LuaState
func Shutdown(st *golua.State) {
	if st != nil {
		st.Close()
	}
}

// Execute executes a simple command, one arg, no results
func Execute(st *golua.State, funcName string, args string) {
	st.GetField(golua.LUA_GLOBALSINDEX, funcName)
	st.PushString(args)
	err := st.Call(1, 0)
	if err != nil {
		log.Println("Lua script error in '", funcName, "' with args '", args, "':", err)
	}
}

// ExecuteInterpret executes a two argument command, passes executing player id
func ExecuteInterpret(st *golua.State, funcName string, actorID string, args string) {
	st.GetField(golua.LUA_GLOBALSINDEX, "execute_character_action")
	st.PushString(actorID)
	st.PushString(funcName)
	st.PushString(args)
	err := st.Call(3, 0)
	if err != nil {
		log.Println("Lua script error in '", funcName, "' with args '", args, "':", err)
	}
}
