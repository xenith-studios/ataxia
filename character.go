package main

type Character struct {
	Id     string
	Player *Player
	Name   string
	World  *World
	Room   *Room
}

func NewCharacter(world *World) *Character {
	ch := Character{
		Id:    uuid(),
		World: world,
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
