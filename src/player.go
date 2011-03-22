/*
    Player structures and functions
*/
package main

import (
    "net/textproto"
    "container/list"
    "os"
    "log"
    "fmt"
    "io"
//    "time"
//    "syscall"
//    "bytes"
    "bufio"
    "ataxia/handler"
)


// The Connection structure wraps all the lower networking details for each connected player
type connection struct {
    socket      io.ReadWriteCloser
    buffer      *bufio.ReadWriter
    server      *Server
    handler     handler.Handler
    remoteAddr  string
    state       string
}


// Account
type Account struct {
    Email string
    Password string
    Name string
    Characters *list.List
}


// Player
type Player struct {
    account Account
    conn *connection
    In chan string
    Out chan string
}


// Player factory
func NewPlayer(conn *connection) (player *Player) {
    player = new(Player)
    player.conn = conn
    player.In = make(chan string, 1024)
    player.Out = make(chan string, 1024)
    player.account.Name = "Unknown"
    return player
}


func (player *Player) Run() {
    buf := make([]byte, 1024)

    // Setup the player here.
    player.Write([]byte("Hello, welcome to Ataxia. What is your account name?\n"))
    if _, err := player.Read(buf); err != nil {
        return
    }
    player.Write([]byte(fmt.Sprintf("Hello %s.\n", string(buf))))
    player.account.Name = string(buf)

    // Create an anonymous goroutine for reading
    go func() {
        for {
            if player.conn.socket == nil {
                break
            }
        
            data := make([]byte, 1024)
            var n int
            var err os.Error
        
            if n, err = player.Read(data); err != nil {
                log.Println(n)
                log.Println(err)
                player.Close()
                break
            }

            if n > 0 {
                player.conn.server.SendToAll(fmt.Sprintf("<%s> %s", player.account.Name, string(data)))
            }
        }
    }()

    // Create an anonymous goroutine for writing
    go func() {
        for {
            if player.conn.socket == nil {
                break
            }
            buf := <-player.In
            if _, err := player.Write([]byte(buf)); err != nil {
                log.Println(err)
                player.Close()
                break
            }
        }
    }()
}


func (player *Player) Close() {
    if (player.conn.socket != nil) {
        player.conn.socket.Close()
        player.conn.socket = nil
        player.conn.buffer = nil
        player.conn.server.RemovePlayer(player)
        log.Println("Player disconnected:", player.account.Name)
    }
}


func (player *Player) Write(buf []byte) (n int, err os.Error) {
    if player.conn.socket == nil || player.conn.buffer == nil {
        return
    }

    if n, err = player.conn.buffer.Write(buf); err != nil {
        if err == os.EOF {
            log.Println("EOF on write, disconnecting player")
            player.Close()
            return 0, nil
        }
        return 0, err
    }
    player.conn.buffer.Flush()
    return
}


func (player *Player) Read(buf []byte) (n int, err os.Error) {
    if player.conn.socket == nil || player.conn.buffer == nil {
        return
    }

    tp := textproto.NewReader(player.conn.buffer.Reader)
    
    var data []byte
    if data, err = tp.ReadLineBytes(); err != nil {
        if err == os.EOF {
            log.Println("Read EOF, disconnecting player")
            player.Close()
            return 0, nil
        }
        return 0, err
    }

    copy(buf, data)
    return len(buf), err
}
