package main

import "ducks/ducks"

func LakeSimulation(ducks []ducks.Duck) {
	for _, duck := range ducks {
		duck.Fly()
		duck.Swim()
	}
}

func main() {
	ducks := []ducks.Duck{
		ducks.NewBlackDuck("Daffy"),
		ducks.NewBlackDuck("Donald"),
		ducks.NewSwan(),
		ducks.NewSwan(),
	}
	LakeSimulation(ducks)
}
