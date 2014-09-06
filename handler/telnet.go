package handler

import (
	"bufio"
	"net"
	//	"log"
	"bytes"
	//	"net/textproto"
	telnet "github.com/xenith-studios/go-telnet"
)

// TelnetHandler type for telnet connections
type TelnetHandler struct {
	buffer *bufio.ReadWriter
	flags  int8
	telnet *telnet.Telnet
}

// NewTelnetHandler returns a new handler
func NewTelnetHandler(conn net.Conn) Handler {
	br := bufio.NewReader(conn)
	bw := bufio.NewWriter(conn)
	return &TelnetHandler{
		buffer: bufio.NewReadWriter(br, bw),
		telnet: telnet.New(),
	}
}

func (handler *TelnetHandler) Read(buf []byte) (n int, err error) {
	var data []byte
	data = make([]byte, 1024)
	if n, err = handler.buffer.Read(data); err != nil {
		return n, err
	}

	// Pass data into telnet.Recv()
	//handler.telnet.Recv(data)

	copy(buf, bytes.Replace(bytes.Replace(data, []byte("\n"), []byte(""), -1), []byte("\r"), []byte(""), -1))
	return n, err
}

func (handler *TelnetHandler) Write(buf []byte) (n int, err error) {
	// Pass the data into telnet.Send()
	//data := handler.telnet.Send(buf)
	data := make([]byte, 4096)
	copy(data, buf)

	if n, err = handler.buffer.Write(data); err != nil {
		return n, err
	}
	handler.buffer.Flush()
	return n, err
}

// Close flushes all remaining data in the buffer and closes everything down
func (handler *TelnetHandler) Close() {
	handler.buffer.Flush()
	handler.buffer = nil
	handler.telnet.Close()
}
