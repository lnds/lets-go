package ducks

import (
	"fmt"
)

type Swan int

func (swan Swan) Fly() {
	fmt.Println("Swan", swan, "is flying")
}

func (swan Swan) Swim() {
	fmt.Println("Swan", swan, "is swimming")
}
