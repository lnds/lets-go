package ducks

import (
	"fmt"
	"math/rand/v2"
)

type Swan struct {
	id int
}

func (swan *Swan) Fly() {
	fmt.Println("Swan", swan.id, "is flying")
}

func (swan *Swan) Swim() {
	fmt.Println("Swan", swan.id, "is swimming")
}

func NewSwan() Duck {
	return &Swan{
		id: rand.IntN(1000),
	}
}
