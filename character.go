package main

import "github.com/xenith-studios/ataxia/utils"

type Character struct {
	Id     string
	Player *Player
	Name   string
	Room   *Room
}

func NewCharacter() *Character {
	ch := Character{
		Id: utils.UUID(),
	}

	return &ch
}
