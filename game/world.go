package game

import (
	"github.com/xenith-studios/ataxia/lua"
	golua "github.com/yuin/gopher-lua"
)

// World defines a single world (an engine can define multiple worlds)
type World struct {
	Interpreter *lua.Interpreter

	// shortcut pointers
	Areas      map[string]*Area
	Characters map[string]*Character
	Rooms      map[string]*Room
	RoomExits  map[string]*RoomExit
}

// NewWorld returns a new World
func NewWorld(state *golua.LState) *World {
	return &World{
		Interpreter: lua.NewInterpreter(state),
		Areas:       make(map[string]*Area),
		Characters:  make(map[string]*Character),
		Rooms:       make(map[string]*Room),
		RoomExits:   make(map[string]*RoomExit),
	}
}

// Initialize initializes a new world and loads all commands into the interpreter
func (world *World) Initialize() {
	world.Interpreter.LoadCommands("scripts/commands/ch_commands.json")

	for _, area := range world.Areas {
		area.Initialize()
	}
}

// LoadAreas loads all areas into the World
func (world *World) LoadAreas() {
	area := NewArea(world)
	area.Load("data/world/midgaard.json")
	world.Areas[area.ID] = area
}

// AddCharacter ##TODO
func (world *World) AddCharacter(ch *Character) {
	ch.World = world
	world.Characters[ch.ID] = ch
}

// AddRoom ##TODO
func (world *World) AddRoom(room *Room) {
	world.Rooms[room.ID] = room
}

// AddRoomExit ##TODO
func (world *World) AddRoomExit(exit *RoomExit) {
	world.RoomExits[exit.ID] = exit
}

// LookupRoom ##TODO
func (world *World) LookupRoom(vnum string) *Room {
	for _, room := range world.Rooms {
		if room.Vnum == vnum {
			return room
		}
	}

	return nil
}

// LoadCharacter loads a character from storage given the character's name
func (world *World) LoadCharacter(name string) *Character {
	// for now, just make a new one, give it a name
	ch := NewCharacter(world)
	ch.Name = name
	ch.Room = world.LookupRoom("3001")
	world.AddCharacter(ch)
	return ch
}
