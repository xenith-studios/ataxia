package handler

type Handler interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Close()
}
