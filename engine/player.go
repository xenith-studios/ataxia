/*
   Player structures and functions
*/
package engine

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
	"github.com/xenith-studios/ataxia/handler"
	"github.com/xenith-studios/ataxia/game"
)

// The Connection structure wraps all the lower networking details for each connected player
type Connection struct {
	socket     io.ReadWriteCloser
	server     *Server
	handler    handler.Handler
	remoteAddr string
	state      string
}

// Account
type Account struct {
	Email      string
	Password   string
	Name       string
	Characters []string
}

// Player
type Player struct {
	account   *Account
	conn      *Connection
	character *game.Character
	In        chan string
	Out       chan string
}

// Player factory
func NewPlayer(server *Server, conn *Connection) (player *Player) {
	player = new(Player)
	player.conn = conn
	player.In = make(chan string, 1024)
	player.Out = make(chan string, 1024)
	player.account = new(Account)
	player.account.Name = "Unknown"
	return player
}

func (player *Player) Run() {
	buf := make([]byte, 1024)

	// Setup the player here.
	player.conn.handler.Write([]byte("Hello, welcome to Ataxia. What is your account name?\n"))
	if _, err := player.conn.handler.Read(buf); err != nil {
		if err == io.EOF {
			log.Println("Read EOF, disconnecting player")
		} else {
			log.Println(err)
		}
		player.Close()
		return
	}
	player.conn.handler.Write([]byte(fmt.Sprintf("Hello %s.\n", string(buf))))
	player.account.Name = string(buf)

	player.character = player.conn.server.World.LoadCharacter(player.account.Name) // let them choose later
	player.character.Attach(player.In)
	player.conn.server.AddPlayer(player)

	// Create an anonymous goroutine for reading
	go func() {
		for {
			if player.conn.socket == nil {
				return
			}

			data := make([]byte, 1024)
			n, err := player.Read(data)

			if err != nil {
				if err == io.EOF {
					log.Println("Read EOF, disconnecting player")
				} else {
					log.Println(err)
				}
				player.Close()
				return
			}

			// TODO: Parse the command here
			if n > 0 {
				data = bytes.Trim(data, " \x00") // trim trailing space and nuls
				player.Interpret(string(data))
				//				player.conn.server.SendToAll(fmt.Sprintf("<%s> %s", player.account.Name, string(data)))
			}
		}
	}()

	// Create an anonymous goroutine for writing
	go func() {
		for line := range player.In {
			if player.conn.socket == nil {
				break
			}
			written := 0
			bytes := []byte(line)
			for written < len(line) {
				n, err := player.Write(bytes[written:])
				if err != nil {
					if err == io.EOF {
						log.Println("EOF on write, disconnecting player")
					} else {
						log.Println(err)
					}
					player.Close()
					return
				}
				written += n
			}
		}
	}()
}

func (player *Player) Close() {
	if player.conn.socket != nil {
		player.conn.socket.Close()
		player.conn.socket = nil
		player.conn.handler.Close()
		player.character.Detach()
		player.conn.server.RemovePlayer(player)
		log.Println("Player disconnected:", player.account.Name)
	}
}

func (player *Player) Interpret(input string) {
	// two level interpeting, do it here (catch player commands), if not found, do it in character

	// interpret goes here

	// else
	player.character.Interpret(input)
	player.character.Write("> ")
}

func (player *Player) Write(buf []byte) (n int, err error) {
	if player.conn.socket == nil {
		return
	}

	return player.conn.handler.Write(buf)
}

func (player *Player) Read(buf []byte) (n int, err error) {
	if player.conn.socket == nil {
		return
	}

	return player.conn.handler.Read(buf)
}
