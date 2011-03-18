package main

import (
    "net"
    "time"
    "container/list"
)

type Connection struct {
    inBuffer []uint8
    outBuffer []uint8
    Socket net.Conn
    Status string
    Account Account
    Player Player
}

type Account struct {
    Email string
    password string
    Characters *list.List
}

type Player struct {
}

var PlayerList *list.List

func init() {
    PlayerList = list.New()
}

func (conn *Connection) Run() {
    for {
        conn.Socket.Write([]uint8("Hello!\n"))
        time.Sleep(60*60*60*60*60)
    }
}