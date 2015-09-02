package lua

import (
	"encoding/json"
	"io/ioutil"
	"log"
	//		"fmt"
	"errors"
	"strings"

	golua "github.com/yuin/gopher-lua"
)

// Command defines a single command from a lua script
type Command struct {
	Script    string
	Func_Name string
	Group     string
}

// Interpreter defines a single command interpreter
type Interpreter struct {
	commandList map[string]Command
	luaState    *golua.LState
}

// NewInterpreter returns a pointer to a new Interpreter
func NewInterpreter(luaState *golua.LState) *Interpreter {
	return &Interpreter{
		luaState: luaState,
		// init stuff
	}
}

// LoadCommands loads all commands from the lua scripts as defined in the commands.json file
func (interp *Interpreter) LoadCommands(commandFile string) {
	bytes, err := ioutil.ReadFile(commandFile)
	if err != nil {
		log.Fatal("Unable to read command list file.")
	}

	err = json.Unmarshal(bytes, &interp.commandList)
	if err != nil {
		log.Fatal("Unable to parse command list.")
	}

	for key := range interp.commandList {
		// need to check and make sure a command with that name was loaded

		// should map these and only try and load lua scripts once, in case multiple commands
		// with same script file
		err := interp.luaState.DoFile(interp.commandList[key].Script)

		if err != nil {
			log.Fatal("Could not read command script,", err) // which one?
			delete(interp.commandList, key)
		}
	}

	log.Printf("Loaded %d commands.", len(interp.commandList))
}

// Interpret does something. ##TODO
func (interp *Interpreter) Interpret(str string, actorID string) error {
	parts := strings.SplitN(str, " ", 2) // need better split (other or multiple whitespace)
	command, found := interp.commandList[parts[0]]

	if !found {
		return errors.New("Command not found")
	}

	args := ""
	if len(parts) > 1 {
		args = strings.TrimSpace(parts[1])
	}

	// acquire lock on player here, to pass UID into lua script.
	ExecuteInterpret(interp.luaState, command.Func_Name, actorID, args)
	return nil
}
