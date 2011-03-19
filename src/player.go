/*
    Player structures and functions
*/
package main

import (
    "net"
//    "time"
    "container/list"
    "os"
    "log"
    "fmt"
    "bytes"
    "bufio"
)

// The Connection structure wraps all the lower networking details for each connected player
type Connection struct {
    socket net.Conn
    reader *bufio.Reader
    State string
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
    connection Connection
    In chan string
    Out chan string
    inBuffer []byte
    outBuffer []byte
}


func (account *Account) Setup() {
//    account.Write([]byte("Hello, welcome to Ataxia. What is your account name?\n"))
//    buf, _ := account.Read()
//    account.Write([]byte(fmt.Sprintf("Hello %s, how are you today?\n", buf)))
}



// Player factory
func NewPlayer(conn net.Conn) (player *Player) {
    player = new(Player)
    player.connection.socket = conn
    player.connection.reader = bufio.NewReader(conn)
    return player
}

func (player *Player) Run() (err os.Error) {
    var data string
    
    // Setup the player here.
    player.Write("Hello, welcome to Ataxia. What is your account name?\n")
    if data, err = player.Read(); err != nil {
        return err
    }
    player.Write(fmt.Sprintf("Hello %s.\n", data))
    player.account.Name = data
    player.Write("\n\n> ")
    //player.connection.socket.SetTimeout(3600)
    for {
        if player.connection.socket == nil {
            break
        }
        
        if data, err = player.Read(); err != nil {
            log.Println(err)
            player.Close()
            break
        }
        
        if len(data) == 0 {
            continue
        }
        
        mainServer.SendToAll(fmt.Sprintf("<%s> %s", player.account.Name, data))
    }
    
    return
}

func (player *Player) Close() {
    if (player.connection.socket != nil) {
        player.connection.socket.Close()
        player.connection.socket = nil
        player.connection.reader = nil
    }
}

func (player *Player) Write(buf string) (n int, err os.Error) {
    n, err = player.connection.socket.Write([]byte(buf))
    return n, err
}

func (player *Player) Read() (data string, err os.Error) {
    buf := make([]byte, 1024)
    if player.connection.reader == nil {
        return
    }

    if buf, err = player.connection.reader.ReadBytes('\n'); err != nil {
        return
    }
    //buf = buf[0:len(buf)-1]
    buf = bytes.TrimSpace(buf)
    buf = bytes.TrimRight(buf, "\n\r")
    data = string(buf)
    log.Println(buf)
    return
}