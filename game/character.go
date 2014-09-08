package game

import (
	"github.com/xenith-studios/ataxia/utils"
)

// Character defines a single character
type Character struct {
	ID     string
	Name   string
	World  *World
	Room   *Room
	output chan string
}

// NewCharacter returns a new charcater
func NewCharacter(world *World) *Character {
	ch := Character{
		World: world,
		ID:    utils.UUID(),
	}

	return &ch
}

// Interpret interprets a single line input from the character
func (ch *Character) Interpret(str string) {
	err := ch.World.Interpreter.Interpret(str, ch.ID)
	if err != nil {
		ch.Write("Huh?\n")
	}
}

// Write to the character
func (ch *Character) Write(str string) {
	if ch.output != nil {
		ch.output <- str
	}
}

// Attach an output channel to the character
func (ch *Character) Attach(c chan string) {
	ch.output = c
}

// Detach the character's output channel
func (ch *Character) Detach() {
	ch.output = nil
}
