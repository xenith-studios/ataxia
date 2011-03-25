/*
    Server structures and functions
*/
package main

import (
		"log"
	"net"
	"os"
	"fmt"
	"bufio"
	"sync"
//	"container/list"
	"ataxia/handler"
)

type PlayerList struct {
	players map[string]*Player
	mu		sync.RWMutex
}


type Server struct {
	socket *net.TCPListener
	PlayerList *PlayerList
	In chan string
	shutdown chan bool
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
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(""), port})
	if err != nil {
		log.Fatalln("Failed to create server:", err)
		return nil
	}

	server = new(Server)
	server.PlayerList = NewPlayerList()
	server.In = make(chan string, 1024)
	server.socket = listener
	server.shutdown = shutdown
	return
}


func (server *Server) Shutdown() {
	if server.socket != nil {
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
			log.Println("Failed to accept new connection")
		} else {
			c := new(Connection)
			c.remoteAddr = conn.RemoteAddr().String()
			c.socket = conn
			c.server = server
			br := bufio.NewReader(conn)
			bw := bufio.NewWriter(conn)
			c.buffer = bufio.NewReadWriter(br, bw)
			c.handler = handler.NewTelnetHandler(conn)
			log.Println("Accepted a new connection:", c.remoteAddr)
			player := NewPlayer(server, c)
			go player.Run()
		}
	}
}


func (server *Server) Run() {
	for {
	}
}


func (server *Server) SendToAll(buf string) {
	for _, player := range server.PlayerList.players {
		if player != nil {
			log.Println(buf)
	        player.In <- fmt.Sprintf("%s\r\n", buf)
		}
    }
}


func (server *Server) AddPlayer(player *Player) {
	server.PlayerList.Add(player.account.Name, player)
}


func (server *Server) RemovePlayer(player *Player) {
    server.PlayerList.Delete(player.account.Name)
}


func (server *Server) Write(buf []byte) (n int, err os.Error) {
	return 0, os.EOF
}
