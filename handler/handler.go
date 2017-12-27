package handler

// Handler defines an interface for input/output handling for each player
type Handler interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Close()
}
