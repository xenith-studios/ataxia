/*
    Server functions
*/
package main

import (
    "log"
    "net"
    "os"
    "fmt"
    "container/list"
    "ataxia/settings"
)

type Server struct {
    Socket *net.TCPListener
    PlayerList *list.List
}

var mainServer Server

func InitializeNetwork() bool {
    if mainServer.Socket != nil {
        return true
    }
    listener, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(""), settings.MainPort})
    if err != nil {
        log.Fatalln("Failed to create server:", err)
        return false
    }
    mainServer.Socket = listener
    mainServer.PlayerList = new(list.List)
    return true
}

func (server Server) Shutdown() {
    if server.Socket != nil {
        for e := server.PlayerList.Front(); e != nil; e = e.Next() {
            player := e.Value.(*Player)
            player.Close()
        }
        server.Socket.Close()
    }
}

func (server Server) Listen() {
    for {
        conn, err := server.Socket.Accept()
        if err != nil {
            log.Println("Failed to accept new connection")
        } else {
            log.Println("Accepted a new connection:", conn.RemoteAddr().String())
            player := NewPlayer(conn)
            server.PlayerList.PushBack(player)
            go player.Run()
        }
    }
}


func (server Server) SendToAll(buf string) {
    for e := server.PlayerList.Front(); e != nil; e = e.Next() {
        player := e.Value.(*Player)
        player.Write(fmt.Sprintf("\n%s\n", buf))
        player.Write("> ")
    }
}


func (server Server) Write(buf []byte) (n int, err os.Error) {
    return 0, os.EOF
}