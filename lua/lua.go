/*
   Lua
*/
package lua

import (
	"log"
	golua "github.com/aarzilli/golua/lua"
//	luar "github.com/stevedonovan/luar"
)

var MainState *golua.State

func NewState() *golua.State {
	log.Println("Initializing Lua State")
	st := golua.NewState()
	st.OpenLibs()
	return st
}

func Shutdown(st *golua.State) {
	if st != nil {
		st.Close()
	}
}

// simple command, one arg, no results
func Execute(st *golua.State, func_name string, args string) {
	st.GetField(golua.LUA_GLOBALSINDEX, func_name)
	st.PushString(args)
	st.Call(1, 0)
}

// two argument execute, passes executing player id
func ExecuteInterpret(st *golua.State, func_name string, player_uid string, args string) {
	st.GetField(golua.LUA_GLOBALSINDEX, func_name)
	st.PushString(player_uid)
	st.PushString(args)
	st.Call(2, 0)
}

