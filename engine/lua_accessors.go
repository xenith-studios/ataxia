package engine

import (
	"fmt"
	"log"

	golua "github.com/yuin/gopher-lua"
	//	"github.com/xenith-studios/ataxia/lua"
	luar "github.com/layeh/gopher-luar"
)

// PublishAccessors registers exported functions into Lua (this is a weird place, should be in main?  or called from there?)
func (server *Server) PublishAccessors(state *golua.LState) {
	state.SetGlobal("GetPlayerData", luar.New(state, server.GetPlayerData))
	state.SetGlobal("SendToPlayers", luar.New(state, server.SendToPlayers))
}

// SendToPlayers sends to all connected players
func (server *Server) SendToPlayers(msg string) {
	for _, player := range server.PlayerList.players {
		if player != nil {
			log.Println(msg)
			player.In <- fmt.Sprintf("%s\r\n", msg)
		}
	}
}

// GetPlayerData returns a single field from a player account structure
func (server *Server) GetPlayerData(id string, field string) string {
	player := server.PlayerList.Get(id)
	var ret string
	if field == "name" { // replace this with reflection on struct tags?
		ret = player.Name
	}
	return ret
}
