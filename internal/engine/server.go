package engine

/*
   Server structures and functions
*/

import (
	"fmt"
	"io"
	"log"
	"net"
	//	"bufio"
	"sync"
	"time"
	//	"container/list"
	"github.com/xenith-studios/ataxia/internal/handler"
)

// PlayerList maintains a list of connected player accounts
type PlayerList struct {
	players map[string]*Account
	mu      sync.RWMutex
}

// Server struct defines main engine data structure
type Server struct {
	socket     *net.TCPListener
	PlayerList *PlayerList
	In         chan string
	shutdown   chan bool
}

// NewPlayerList returns a new PlayerList struct
func NewPlayerList() *PlayerList {
	return &PlayerList{players: make(map[string]*Account)}
}

// Add a new account to the PlayerList
func (list *PlayerList) Add(name string, player *Account) {
	list.mu.Lock()
	defer list.mu.Unlock()
	list.players[name] = player
}

// Delete an account from the PlayerList
func (list *PlayerList) Delete(name string) {
	list.mu.Lock()
	defer list.mu.Unlock()
	list.players[name] = nil
}

// Get and return a player account by exact name
func (list *PlayerList) Get(name string) *Account {
	list.mu.RLock()
	defer list.mu.RUnlock()
	return list.players[name]
}

// NewServer creates a new server and returns a pointer to it
func NewServer(port int, shutdown chan bool) *Server {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP(""), Port: port, Zone: ""})
	if err != nil {
		log.Fatalln("Failed to create server:", err)
		return nil
	}

	server := &Server{
		PlayerList: NewPlayerList(),
		socket:     listener,
		shutdown:   shutdown,
		In:         make(chan string, 1024),
	}

	return server
}

// Shutdown the server and disconnect all remaining player accounts
func (server *Server) Shutdown() {
	if server.socket != nil {
		server.SendToPlayers("Server is shutting down!")
		for _, player := range server.PlayerList.players {
			if player != nil {
				player.Close()
			}
		}
		server.socket.Close()
		server.socket = nil
	}
}

// Listen is the goroutine that accepts new player connections and creates the account structure
func (server *Server) Listen() {
	for {
		if server.socket == nil {
			log.Println("Server socket closed")
			server.shutdown <- true
			return
		}
		conn, err := server.socket.Accept()
		if err != nil {
			log.Println("Failed to accept new connection:", err)
			continue
		} else {
			c := new(connection)
			c.remoteAddr = conn.RemoteAddr().String()
			c.socket = conn
			c.handler = handler.NewTelnetHandler(conn)
			log.Println("Accepted a new connection:", c.remoteAddr)
			player := NewAccount(server, c)
			go player.Run()
		}
	}
}

// Run is the main goroutine for the engine. It handles all game updates and events, also known as the main loop.
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

// AddPlayer adds a new player to the PlayerList
func (server *Server) AddPlayer(player *Account) {
	server.PlayerList.Add(player.Name, player)
}

// RemovePlayer removes a player from the PlayerList
func (server *Server) RemovePlayer(player *Account) {
	server.PlayerList.Delete(player.Name)
}

// Write is a convenience function to satisfy the io.Writer interface
func (server *Server) Write(buf []byte) (int, error) {
	return 0, io.EOF
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
