package engine

/*
   Account structures and functions
*/

import (
	//	"net/textproto"
	//	"container/list"
	"bytes"
	"fmt"
	"io"
	"log"
	//	"time"
	//	"syscall"
	//	"bytes"
	//	"bufio"
	//	"strings"
	"github.com/xenith-studios/ataxia/game"
	"github.com/xenith-studios/ataxia/handler"
)

// The Account struct defines each connected player at the engine level
type Account struct {
	Email      string
	Password   string
	Name       string
	Characters []string
	conn       *connection
	server     *Server
	character  *game.Character
	In         chan string
	Out        chan string
}

// The connection struct wraps all the lower-level networking details for each connected player
type connection struct {
	socket     io.ReadWriteCloser
	handler    handler.Handler
	remoteAddr string
	state      string
}

// Account factory
func NewAccount(server *Server, conn *connection) (account *Account) {
	account = new(Account)
	account.conn = conn
	account.server = server
	account.In = make(chan string, 1024)
	account.Out = make(chan string, 1024)
	account.Name = "Unknown"
	return account
}

func (account *Account) Run() {
	buf := make([]byte, 1024)

	// Setup the account here.
	account.conn.handler.Write([]byte("Hello, welcome to Ataxia. What is your account name?\n"))
	if _, err := account.conn.handler.Read(buf); err != nil {
		if err == io.EOF {
			log.Println("Read EOF, disconnecting player...")
		} else {
			log.Println(err)
		}
		account.Close()
		return
	}
	account.conn.handler.Write([]byte(fmt.Sprintf("Hello %s.\n", string(buf))))
	account.Name = string(buf)

	account.character = account.server.World.LoadCharacter(account.Name) // let them choose later
	account.character.Attach(account.In)
	account.server.AddPlayer(account)

	// Create an anonymous goroutine for reading
	go func() {
		for {
			if account.conn.socket == nil {
				return
			}

			data := make([]byte, 1024)
			n, err := account.Read(data)

			if err != nil {
				if err == io.EOF {
					log.Println("Read EOF, disconnecting account")
				} else {
					log.Println(err)
				}
				account.Close()
				return
			}

			// TODO: Parse the command here
			if n > 0 {
				data = bytes.Trim(data, " \x00") // trim trailing space and nuls
				account.Interpret(string(data))
				//				account.server.SendToAll(fmt.Sprintf("<%s> %s", account.Name, string(data)))
			}
		}
	}()

	// Create an anonymous goroutine for writing
	go func() {
		for line := range account.In {
			if account.conn.socket == nil {
				break
			}
			written := 0
			bytes := []byte(line)
			for written < len(line) {
				n, err := account.Write(bytes[written:])
				if err != nil {
					if err == io.EOF {
						log.Println("EOF on write, disconnecting account")
					} else {
						log.Println(err)
					}
					account.Close()
					return
				}
				written += n
			}
		}
	}()
}

func (account *Account) Close() {
	if account.conn.socket != nil {
		account.conn.handler.Close()
		account.character.Detach()
		account.server.RemovePlayer(account)
		account.conn.socket.Close()
		account.conn.socket = nil
		log.Println("Account disconnected:", account.Name)
	}
}

func (account *Account) Interpret(input string) {
	// two level interpeting, do it here (catch account commands), if not found, do it in character

	// interpret goes here

	// else
	account.character.Interpret(input)
	account.character.Write("> ")
}

func (account *Account) Write(buf []byte) (n int, err error) {
	if account.conn.socket == nil {
		return
	}

	return account.conn.handler.Write(buf)
}

func (account *Account) Read(buf []byte) (n int, err error) {
	if account.conn.socket == nil {
		return
	}

	return account.conn.handler.Read(buf)
}
