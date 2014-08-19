package game

import (
	golua "github.com/aarzilli/golua/lua"
	"github.com/xenith-studios/ataxia/lua"

)

type World struct {
	Interpreter *lua.Interpreter

	// shortcut pointers
	Areas      map[string]*Area
	Characters map[string]*Character
	Rooms      map[string]*Room
	RoomExits  map[string]*RoomExit
}

func NewWorld(state *golua.State) *World {
	world := World{
		Interpreter: lua.NewInterpreter(state),
		Areas:       make(map[string]*Area),
		Characters:  make(map[string]*Character),
		Rooms:       make(map[string]*Room),
		RoomExits:   make(map[string]*RoomExit),
	}
	return &world
}

func (world *World) Initialize() {
	world.Interpreter.LoadCommands("scripts/commands/ch_commands.json")

	for _, area := range world.Areas {
		area.Initialize()
	}
}

func (world *World) LoadAreas() {
	area := NewArea(world)
	area.Load("data/world/midgaard.json")
	world.Areas[area.ID] = area
}

func (world *World) AddCharacter(ch *Character) {
	ch.World = world
	world.Characters[ch.Id] = ch
}

func (world *World) AddRoom(room *Room) {
	world.Rooms[room.Id] = room
}

func (world *World) AddRoomExit(exit *RoomExit) {
	world.RoomExits[exit.ID] = exit
}

func (world *World) LookupRoom(vnum string) *Room {
	for _, room := range world.Rooms {
		if room.Vnum == vnum {
			return room
		}
	}

	return nil
}

func (world *World) LoadCharacter(name string) (ch *Character) {
	// for now, just make a new one, give it a name
	ch = NewCharacter(world)
	ch.Name = name
	ch.Room = world.LookupRoom("3001")
	world.AddCharacter(ch)
	return
}
