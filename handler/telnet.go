package handler

import (
	"bufio"
	"net"
	//	"log"
	"bytes"
	//	"net/textproto"
	"github.com/xenith-studios/ataxia/telnet"
)

type TelnetHandler struct {
	buffer *bufio.ReadWriter
	flags  int8
	telnet *telnet.Telnet
}

func NewTelnetHandler(conn net.Conn) (handler *TelnetHandler) {
	handler = new(TelnetHandler)
	br := bufio.NewReader(conn)
	bw := bufio.NewWriter(conn)
	handler.buffer = bufio.NewReadWriter(br, bw)
	handler.telnet = telnet.New()
	return
}

func (handler *TelnetHandler) Read(buf []byte) (n int, err error) {
	var data []byte
	data = make([]byte, 1024)
	if n, err = handler.buffer.Read(data); err != nil {
		return n, err
	}

	// Pass data into telnet.Recv()
	handler.telnet.Recv(data)

	copy(buf, bytes.Replace(bytes.Replace(data, []byte("\n"), []byte(""), -1), []byte("\r"), []byte(""), -1))
	return n, err
}

func (handler *TelnetHandler) Write(buf []byte) (n int, err error) {
	// Pass the data into telnet.Send()
	data := handler.telnet.Send(buf)

	if n, err = handler.buffer.Write(data); err != nil {
		return n, err
	}
	handler.buffer.Flush()
	return n, err
}

func (handler *TelnetHandler) Close() {
	handler.buffer = nil
	handler.telnet.Close()
}
