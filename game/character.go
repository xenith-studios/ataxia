package game

import (
	"github.com/xenith-studios/ataxia/utils"
)

type Character struct {
	Id     string
	Player *Player
	Name   string
	World  *World
	Room   *Room
}

func NewCharacter(world *World) *Character {
	ch := Character{
		World: world,
		Id:    utils.UUID(),
	}

	return &ch
}

func (ch *Character) Interpret(str string) {
	ch.World.Interpreter.Interpret(str, ch)
}

func (ch *Character) Write(str string) {
	if ch.Player != nil {
		ch.Player.Write([]byte(str))
	}
}
