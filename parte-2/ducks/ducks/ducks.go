package ducks

import "fmt"

type Duck interface {
	Fly()
	Swim()
}

type BlackDuck struct {
	name string
}

func (duck *BlackDuck) Fly() {
	fmt.Printf("%s duck is flying\n", duck.name)
}

func (duck *BlackDuck) Swim() {
	fmt.Printf("%s duck is swimming\n", duck.name)
}

func NewBlackDuck(name string) Duck {
	return &BlackDuck{
		name: name,
	}
}
