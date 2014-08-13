/*
   Server structures and functions
*/
package main

import (
	"io"
	"log"
	"net"
	//	"bufio"
	"sync"
	"time"
	//	"container/list"
	golua "github.com/aarzilli/golua/lua"
	"github.com/xenith-studios/ataxia/handler"
	"github.com/xenith-studios/ataxia/lua"
)

type PlayerList struct {
	players map[string]*Player
	mu      sync.RWMutex
}

type Server struct {
	socket     *net.TCPListener
	luaState   *golua.State
	World      *World
	PlayerList *PlayerList
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

func NewServer(port int, shutdown chan bool) (server *Server) {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP(""), Port: port, Zone: ""})
	if err != nil {
		log.Fatalln("Failed to create server:", err)
		return nil
	}

	server = new(Server)

	server.luaState = lua.MainState
	server.PublishAccessors(server.luaState)

	server.World = NewWorld(server.luaState)
	server.World.PublishAccessors(server.luaState)

	server.PlayerList = NewPlayerList()

	server.In = make(chan string, 1024)
	server.socket = listener
	server.shutdown = shutdown
	return
}

func (server *Server) InitializeWorld() {
	server.World.LoadAreas()
	server.World.Initialize()

}

func (server *Server) Shutdown() {
	if server.socket != nil {
		server.SendToPlayers("Server is shutting down!")
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

func (server *Server) AddPlayer(player *Player) {
	server.PlayerList.Add(player.account.Name, player)
}

func (server *Server) RemovePlayer(player *Player) {
	server.PlayerList.Delete(player.account.Name)
}

func (server *Server) Write(buf []byte) (n int, err error) {
	return 0, io.EOF
}
