package main

import (
	"fmt"
	golua "github.com/aarzilli/golua/lua"
	"log"
	//	"github.com/xenith-studios/ataxia/lua"
	luar "github.com/stevedonovan/luar"
)

func (server *Server) PublishAccessors(state *golua.State) {
	// register exported functions (this is a weird place, should be in main?  or called from there?)
	luar.Register(state, "", luar.Map{
		"GetPlayerData": server.GetPlayerData,
		"SendToPlayers": server.SendToPlayers,
	})
}

func (server *Server) SendToPlayers(msg string) {
	for _, player := range server.PlayerList.players {
		if player != nil {
			log.Println(msg)
			player.In <- fmt.Sprintf("%s\r\n", msg)
		}
	}
}

func (server *Server) GetPlayerData(id string, field string) (ret string) {
	player := server.PlayerList.Get(id)
	if field == "name" { // replace this with reflection on struct tags?
		ret = player.account.Name
	}
	return
}

func (world *World) PublishAccessors(state *golua.State) {
	luar.Register(state, "", luar.Map{
		"SendToAll":        world.SendToAll,
		"SendToOthers":     world.SendToOthers,
		"SendToChar":       world.SendToChar,
		"GetCharacterData": world.GetCharacterData,
		"GetRoomData":      world.GetRoomData,
	})
}

func (world *World) SendToAll(msg string) {
	for _, ch := range world.Characters {
		if ch.Player != nil {
			log.Println(msg)
			ch.Player.In <- fmt.Sprintf("%s\r\n", msg)
		}
	}
}

func (world *World) SendToOthers(char_id string, msg string) {
	for id, ch := range world.Characters {
		if id == char_id {
			continue
		}

		if ch.Player != nil {
			log.Println(msg)
			ch.Player.In <- fmt.Sprintf("%s\r\n", msg)
		}
	}
}

func (world *World) SendToChar(id string, msg string) {
	ch := world.Characters[id]
	if ch != nil {
		if ch.Player != nil {
			ch.Player.In <- msg
		}
	}
}

func (world *World) GetCharacterData(id string, field string) (ret string) {
	ch := world.Characters[id]
	if ch == nil {
		return ""
	}

	if field == "name" {
		return ch.Name
	}
	if field == "room" {
		return ch.Room.Id
	}
	return
}

func (world *World) GetRoomData(id string, field string) (ret string) {
	ch := world.Rooms[id]
	if ch == nil {
		return ""
	}

	if field == "name" {
		return ch.Name
	}
	if field == "description" {
		return ch.Description
	}
	return
}
