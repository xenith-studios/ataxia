/*
   Command interpreter
*/
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
//	"fmt"
	"strings"
	"github.com/xenith-studios/ataxia/lua"
)

type Command struct {
	Script		string
	Func_name	string
	Group		string
}

var  commandList map[string]Command // pointers so they can be easily modified below

func LoadCommandList() {
	bytes, err := ioutil.ReadFile("scripts/commands/commands.json")
	if err != nil {
		log.Fatal("Unable to read command list file.")
	}

	err = json.Unmarshal(bytes, &commandList)
	if err != nil {
		log.Fatal("Unable to parse command list.")
	}

	L := lua.MainState

	for key := range commandList {
		// need to check and make sure a command with that name was loaded

		// should map these and only try and load lua scripts once, in case multiple commands
		// with same script file
		err := L.DoFile(commandList[key].Script)

		if err != nil {
			log.Fatal("Could not read command script,", err) // which one?
			delete(commandList, key)
		}
	}

	log.Printf("Loaded %d commands.", len(commandList))
}

func Interpret(str string, player *Player) {
	parts := strings.SplitN(str, " ", 2) // need better split (other or multiple whitespace)
	command, found := commandList[parts[0]]

	if !found {
		player.Write([]byte("Huh?\n"))
	}

	args := ""
	if len(parts) > 1 {
		args = parts[1]
	}

	lua.Execute(command.Func_name, args)
}
