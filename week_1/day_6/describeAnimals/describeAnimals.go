package main

import (
	"fmt"
)

type Flyer interface {
	Fly() string
}

type Swimmer interface {
	Swim() string
}

type Duck struct {
	Name string
}

func (d Duck) Fly() string {
	return fmt.Sprintf("My name is %s and I'm flying!", d.Name)
}

func (d Duck) Swim() string {
	return fmt.Sprintf("My name is %s and I'm swimming", d.Name)
}

func describeAnimals(animal interface{}) {
	var canSwim string
	var canFly string
	if swimmer, ok := animal.(Swimmer); ok {
		canSwim = swimmer.Swim()
	} else {
		canSwim = ""
	}

	if flyer, ok := animal.(Flyer); ok {
		canFly = flyer.Fly()
	} else {
		canFly = ""
	}

	fmt.Println("This animal can:\n\tFly:", canFly, "\n\tSwim:", canSwim)
}

func main() {
	duck := Duck{Name: "Waffle"}
	describeAnimals(duck)
}
