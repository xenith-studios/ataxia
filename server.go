/*
   Server structures and functions
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	//	"bufio"
	"sync"
	"time"
	//	"container/list"
	"github.com/xenith-studios/ataxia/handler"
)

type PlayerList struct {
	players map[string]*Player
	mu      sync.RWMutex
}

type Server struct {
	socket     *net.TCPListener
	PlayerList *PlayerList
	AreaList	[]*Area
	CharacterList	map[string]*Character
	RoomList		map[string]*Room
	In         chan string
	shutdown   chan bool
}

func NewPlayerList() (list *PlayerList) {
	return &PlayerList{players: make(map[string]*Player)}
}

func (list *PlayerList) Add(name string, player *Player) {
	list.mu.Lock()
	defer list.mu.Unlock()
	list.players[name] = player
}

func (list *PlayerList) Delete(name string) {
	list.mu.Lock()
	defer list.mu.Unlock()
	list.players[name] = nil
}

func (list *PlayerList) Get(name string) (player *Player) {
	list.mu.RLock()
	defer list.mu.RUnlock()
	player = list.players[name]
	return
}


func (server *Server) LoadAreas() {
	area := NewArea()
	area.Server = server
	area.Load("data/world/midgaard.json")
	server.AreaList = []*Area{area}
}

func (server *Server) InitializeWorld() {
	for area := range server.AreaList {
		server.AreaList[area].Initialize()
	}
}

func NewServer(port int, shutdown chan bool) (server *Server) {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(""), port, ""})
	if err != nil {
		log.Fatalln("Failed to create server:", err)
		return nil
	}

	server = new(Server)
	server.PlayerList = NewPlayerList()
	server.CharacterList = make(map[string]*Character)
	server.RoomList = make(map[string]*Room)
	server.In = make(chan string, 1024)
	server.socket = listener
	server.shutdown = shutdown
	return
}

func (server *Server) Shutdown() {
	if server.socket != nil {
		server.SendToAll("Server is shutting down!")
		for _, player := range server.PlayerList.players {
			if player != nil {
				player.Close()
			}
		}
		server.socket.Close()
	}
}

func (server *Server) Listen() {
	for {
		if server.socket == nil {
			log.Println("Server socket closed")
			shutdown <- true
			return
		}
		conn, err := server.socket.Accept()
		if err != nil {
			log.Println("Failed to accept new connection:", err)
		} else {
			c := new(Connection)
			c.remoteAddr = conn.RemoteAddr().String()
			c.socket = conn
			c.server = server
			c.handler = handler.NewTelnetHandler(conn)
			log.Println("Accepted a new connection:", c.remoteAddr)
			player := NewPlayer(server, c)
			go player.Run()
		}
	}
}

func (server *Server) Run() {
	// Main loop
	// Handle network messages (push user events)
	// Handle game updates
	// Game tick
	// Time update
	// Weather update
	// Entity updates (push events)
	// Handle pending events
	// Handle pending messages (network and player)
	// Sleep
	for {
		// Sleep for 1 ms
		time.Sleep(1000000)
	}
}

func (server *Server) SendToAll(msg string) {
	for _, ch := range server.CharacterList {
		if ch.Player != nil {
			log.Println(msg)
			ch.Player.In <- fmt.Sprintf("%s\r\n", msg)
		}
	}
}

func (server *Server) SendToOthers(char_id string, msg string) {
	for id, ch := range server.CharacterList {
		if id == char_id {
			continue
		}

		if ch.Player != nil {
			log.Println(msg)
			ch.Player.In <- fmt.Sprintf("%s\r\n", msg)
		}
	}
}

func (server *Server) SendToChar(id string, msg string) {
	ch := server.CharacterList[id]
	if ch != nil {
		if ch.Player != nil {
			ch.Player.In <- msg
		}
	}
}

// for exporting to lua
func (server *Server) GetPlayerData(id string, field string) (ret string) {
	player := server.PlayerList.Get(id)
	if field == "name" {  // replace this with reflection on struct tags?
		ret = player.account.Name
	}
	return
}

func (server *Server) GetCharacterData(id string, field string) (ret string) {
	ch := server.CharacterList[id]
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

func (server *Server) GetRoomData(id string, field string) (ret string) {
	ch := server.RoomList[id]
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


func (server *Server) AddPlayer(player *Player) {
	server.PlayerList.Add(player.account.Name, player)
}

func (server *Server) AddCharacter(ch *Character) {
	server.CharacterList[ch.Id] = ch
}

func (server *Server) AddRoom(room *Room) {
	server.RoomList[room.Id] = room
}

func (server *Server) RemovePlayer(player *Player) {
	server.PlayerList.Delete(player.account.Name)
}

func (server *Server) Write(buf []byte) (n int, err error) {
	return 0, io.EOF
}
