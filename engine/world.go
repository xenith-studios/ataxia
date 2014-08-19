package engine

import (
	golua "github.com/aarzilli/golua/lua"
	"github.com/xenith-studios/ataxia/game"
	"github.com/xenith-studios/ataxia/interpret"
)

type World struct {
	Interpreter *interpet.Interpreter

	// shortcut pointers
	Areas      map[string]*game.Area
	Characters map[string]*game.Character
	Rooms      map[string]*game.Room
	RoomExits  map[string]*game.RoomExit
}

func NewWorld(state *golua.State) *World {
	world := World{
		Interpreter: NewInterpreter(state),
		Areas:       make(map[string]*game.Area),
		Characters:  make(map[string]*game.Character),
		Rooms:       make(map[string]*game.Room),
		RoomExits:   make(map[string]*game.RoomExit),
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
	area := game.NewArea(world)
	area.Load("data/world/midgaard.json")
	world.Areas[area.ID] = area
}

func (world *World) AddCharacter(ch *game.Character) {
	ch.World = world
	world.Characters[ch.Id] = ch
}

func (world *World) AddRoom(room *game.Room) {
	world.Rooms[room.Id] = room
}

func (world *World) AddRoomExit(exit *game.RoomExit) {
	world.RoomExits[exit.ID] = exit
}

func (world *World) LookupRoom(vnum string) *game.Room {
	for _, room := range world.Rooms {
		if room.Vnum == vnum {
			return room
		}
	}

	return nil
}

func (world *World) LoadCharacter(name string) (ch *game.Character) {
	// for now, just make a new one, give it a name
	ch = game.NewCharacter(world)
	ch.Name = name
	ch.Room = world.LookupRoom("3001")
	world.AddCharacter(ch)
	return
}
