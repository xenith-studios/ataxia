/*
    Net
*/
package net

import (
    "log"
    "net"
    "fmt"
    "ataxia/settings"
)

var Server *net.Listener

func Initialize() {
    server, err := net.Listen("tcp", fmt.Sprintf(":%d", settings.MainPort))
    if err != nil {
        log.Fatalln("Failed to create server:", err)
    }
    Server = &server
}