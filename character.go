package main

type Character struct {
	Id			string
	Player		*Player
	Name		string
	Room		*Room
}

func NewCharacter() *Character {
	ch := Character {
		Id:		uuid(),
	}

	return &ch
}
