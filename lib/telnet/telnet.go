package telnet

// #include "telnet.h"
import "C"

type Telnet_T struct{
	
}

func NewTelnet() *Telnet_T {
	telnet := new(Telnet_T)
	return telnet
}