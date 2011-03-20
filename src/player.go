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
//    "time"
//    "syscall"
//    "bytes"
//    "bufio"
)


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
                mainServer.SendToAll(fmt.Sprintf("<%s> %s", player.account.Name, string(data)))
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
        log.Println("Player disconnected:", player.account.Name)
    }
}

func (player *Player) Write(buf []byte) (n int, err os.Error) {
    n, err = player.conn.buffer.Write(buf)
    player.conn.buffer.Flush()
    return
}

func (player *Player) Read(buf []byte) (n int, err os.Error) {
    if player.conn.buffer == nil {
        return
    }
    
    tp := textproto.NewReader(player.conn.buffer.Reader)
    
    var s string
    if s, err = tp.ReadLine(); err != nil {
        if err == os.EOF {
            log.Println("Read EOF, disconnecting player")
            player.Close()
            return 0, nil
        }
        return 0, err
    }

    copy(buf, s)
    return len(s), err
}
