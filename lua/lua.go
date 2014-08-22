/*
   Lua
*/
package lua

import (
	golua "github.com/aarzilli/golua/lua"
	"log"
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
	err := st.Call(1, 0)
	if err != nil {
		log.Println("Lua script error in '", func_name, "' with args '", args, "':", err)
	}
}

// two argument execute, passes executing player id
func ExecuteInterpret(st *golua.State, func_name string, actor_id string, args string) {
	st.GetField(golua.LUA_GLOBALSINDEX, "execute_character_action")
	st.PushString(actor_id)
	st.PushString(func_name)
	st.PushString(args)
	err := st.Call(3, 0)
	if err != nil {
		log.Println("Lua script error in '", func_name, "' with args '", args, "':", err)
	}
}
