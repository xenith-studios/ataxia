package handler

import (
	"os"
	"bufio"
	"net"
//	"log"
	"net/textproto"
)


type TelnetHandler struct {
	buffer *bufio.ReadWriter
	flags int8
}


func NewTelnetHandler(conn net.Conn) (handler *TelnetHandler) {
	handler = new(TelnetHandler)
	br := bufio.NewReader(conn)
	bw := bufio.NewWriter(conn)
	handler.buffer = bufio.NewReadWriter(br, bw)
	return
}


func (handler *TelnetHandler) Read(buf []byte) (n int, err os.Error) {
	tp := textproto.NewReader(handler.buffer.Reader)

	var data []byte
	if data, err = tp.ReadLineBytes(); err != nil {
		return len(data), err
	}

	// Pass data into telnet.Recv()
	//telnet.Recv()
	copy(buf, data)
	return len(data), err
}


func (handler *TelnetHandler) Write(buf []byte) (n int, err os.Error) {
	if n, err = handler.buffer.Write(buf); err != nil {
		return n, err
	}
	handler.buffer.Flush()
	return n, err
}


func (handler *TelnetHandler) Close() {
	handler.buffer = nil
}