/*
    Net
*/
package server

import (
    "log"
    "net"
    "ataxia/game"
    "ataxia/settings"
)

var (
    Server *net.TCPListener
    Descriptor int
)


func Initialize() bool {
    if Server != nil {
        return true
    }
    listener, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(""), settings.MainPort})
    if err != nil {
        log.Fatalln("Failed to create server:", err)
    }
    Server = listener
    return true
}

func Shutdown() {
    if Server != nil {
        Server.Close()
    }
}

func Listen() {
    for {
        conn, err := Server.Accept()
        if err != nil {
            log.Println("Failed to accept a connection")
        } else {
            log.Println("Accepted a connection")
            player := new(game.Player)
            player.Socket = conn
        }
        conn.Close()
    }
}