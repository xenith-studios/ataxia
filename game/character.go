package game

import (
	"github.com/xenith-studios/ataxia/utils"
)

type Character struct {
	Id     string
	Name   string
	World  *World
	Room   *Room
	output chan string
}

func NewCharacter(world *World) *Character {
	ch := Character{
		World: world,
		Id:    utils.UUID(),
	}

	return &ch
}

func (ch *Character) Interpret(str string) {
	err := ch.World.Interpreter.Interpret(str, ch.Id)
	if err != nil {
		ch.Write("Huh?\n")
	}
}

func (ch *Character) Write(str string) {
	if ch.output != nil {
		ch.output <- str
	}
}

func (ch *Character) Attach(c chan string) {
	ch.output = c
}

func (ch *Character) Detach() {
	ch.output = nil
}