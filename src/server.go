/*
    Networking functions
*/
package main

import (
    "log"
    "net"
    "ataxia/settings"
)

type Server struct {
    Socket *net.TCPListener
}

var mainServer *Server

func InitializeNetwork() bool {
    if mainServer != nil {
        return true
    }
    mainServer = new(Server)
    listener, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(""), settings.MainPort})
    if err != nil {
        log.Fatalln("Failed to create server:", err)
        return false
    }
    mainServer.Socket = listener
    return true
}

func (server *Server) Shutdown() {
    if server.Socket != nil {
        server.Socket.Close()
    }
}

func (server *Server) Listen() {
    for {
        conn, err := server.Socket.Accept()
        if err != nil {
            log.Println("Failed to accept new connection")
        } else {
            log.Println("Accepted a new connection:", conn.RemoteAddr().String())
            connection := new(Connection)
            connection.Socket = conn
            connection.Status = "connecting"
            PlayerList.PushBack(connection)
            go connection.Run()
        }
    }
}