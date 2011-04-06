package handler

import (
	"os"
	"bufio"
	"net"
//	"log"
	"bytes"
//	"net/textproto"
	"telnet"
)


type TelnetHandler struct {
	buffer *bufio.ReadWriter
	flags int8
	telnet_t *telnet.Telnet_T
}


func NewTelnetHandler(conn net.Conn) (handler *TelnetHandler) {
	handler = new(TelnetHandler)
	br := bufio.NewReader(conn)
	bw := bufio.NewWriter(conn)
	handler.buffer = bufio.NewReadWriter(br, bw)
	handler.telnet_t = telnet.NewTelnet()
	return
}


func (handler *TelnetHandler) Read(buf []byte) (n int, err os.Error) {
	var data []byte
	data = make([]byte, 1024)
	if n, err = handler.buffer.Read(data); err != nil {
		return n, err
	}

	// Pass data into telnet.Recv()
	//telnet.Recv()

	copy(buf, bytes.Replace(bytes.Replace(data, []byte("\n"), []byte(""), -1), []byte("\r"), []byte(""), -1))
	return n, err
}


func (handler *TelnetHandler) Write(buf []byte) (n int, err os.Error) {
	// Pass the data into telnet.Send()
	//telet.Send()

	if n, err = handler.buffer.Write(buf); err != nil {
		return n, err
	}
	handler.buffer.Flush()
	return n, err
}


func (handler *TelnetHandler) Close() {
	handler.buffer = nil
}
