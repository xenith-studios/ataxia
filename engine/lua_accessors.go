package engine

import (
	"fmt"
	"log"

	golua "github.com/aarzilli/golua/lua"
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
		ret = player.Name
	}
	return
}
