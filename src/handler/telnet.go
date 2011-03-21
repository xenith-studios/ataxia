package handler

import (
    "os"
)


type TelnetHandler struct {
    
}


func NewTelnetHandler() (handler *TelnetHandler) {
    handler = new(TelnetHandler)
    return
}


func (handler *TelnetHandler) Read(buf []byte) (n int, err os.Error) {
    return 0, nil
}