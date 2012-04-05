package telnet

/*
#include <stdlib.h>
#include "telnet.h"
*/
import "C"

type Telnet struct{
	telnet_t *C.telnet_t
}

func NewTelnet() *Telnet {
	telnet := new(Telnet)
	telnet.telnet_t = C.initialize()
	return telnet
}

func (telnet *Telnet) Close() {
	C.cleanup(telnet.telnet_t)
	telnet.telnet_t = nil
}

func (telnet *Telnet) Send(buf []byte) (data []byte) {
	data = make([]byte, 4096)
	copy(data, buf)
	return data
}

func (telnet *Telnet) Recv(buf []byte) {
	C.telnet_recv(telnet.telnet_t, C.CString(string(buf)), C.size_t(len(buf)))
}