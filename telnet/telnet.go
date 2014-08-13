package telnet

type Telnet struct {
}

func New() *Telnet {
	telnet := new(Telnet)
	return telnet
}

func (telnet *Telnet) Close() {
}

func (telnet *Telnet) Send(buf []byte) (data []byte) {
	data = make([]byte, 4096)
	copy(data, buf)
	return data
}

func (telnet *Telnet) Recv(buf []byte) {
}
