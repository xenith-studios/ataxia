/*
    Server functions
*/
package main

import (
    "log"
    "net"
    "os"
    "fmt"
    "io"
    "bufio"
    "container/list"
    "ataxia/settings"
)

type Server struct {
    socket *net.TCPListener
    PlayerList *list.List
    In chan string
}

// The Connection structure wraps all the lower networking details for each connected player
type connection struct {
    socket      io.ReadWriteCloser
    buffer      *bufio.ReadWriter
    remoteAddr  string
    state       string
}

var mainServer *Server


func NewServer() (server *Server) {
    listener, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(""), settings.MainPort})
    if err != nil {
        log.Fatalln("Failed to create server:", err)
        return nil
    }

    server = new(Server)
    server.PlayerList = new(list.List)
    server.In = make(chan string, 1024)
    server.socket = listener
    return
}


func (server *Server) Shutdown() {
    if server.socket != nil {
        for e := server.PlayerList.Front(); e != nil; e = e.Next() {
            player := e.Value.(*Player)
            player.Close()
        }
        server.socket.Close()
    }
}

func (server *Server) Listen() {
    for {
        conn, err := server.socket.Accept()
        if err != nil {
            log.Println("Failed to accept new connection")
        } else {
            c := new(connection)
            c.remoteAddr = conn.RemoteAddr().String()
            c.socket = conn
            br := bufio.NewReader(conn)
            bw := bufio.NewWriter(conn)
            c.buffer = bufio.NewReadWriter(br, bw)
            log.Println("Accepted a new connection:", c.remoteAddr)
            player := NewPlayer(c)
            server.PlayerList.PushBack(player)
            go player.Run()
        }
    }
}


func (server *Server) Run() {
    for {
    }
}

func (server *Server) SendToAll(buf string) {
    for e := server.PlayerList.Front(); e != nil; e = e.Next() {
        player := e.Value.(*Player)
        log.Println(buf)
        player.In <- fmt.Sprintf("\n\r%s\n\r", buf)
    }
}


func (server *Server) Write(buf []byte) (n int, err os.Error) {
    return 0, os.EOF
}