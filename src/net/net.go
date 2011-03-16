/*
    Net
*/
package net

import (
    "log"
    "net"
    //"fmt"
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
