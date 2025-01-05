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

type Fish struct {
	Name string
}

type Owl struct {
	Name string
}

func (o Owl) Fly() string {
	return fmt.Sprintf("My name is %s and I'm flying", o.Name)
}

func (f Fish) Swim() string {
	return fmt.Sprintf("My name is %s and I'm swimming", f.Name)
}

func (d Duck) Fly() string {
	return fmt.Sprintf("My name is %s and I'm flying!", d.Name)
}

func (d Duck) Swim() string {
	return fmt.Sprintf("My name is %s and I'm swimming", d.Name)
}

func describeAnimals(animal interface{}) {
	fmt.Println("This animal can:")
	if swimmer, ok := animal.(Swimmer); ok {
		fmt.Println("\t", swimmer.Swim())
	}
	if flyer, ok := animal.(Flyer); ok {
		fmt.Println("\t", flyer.Fly())
	}
}

func main() {
	animals := []interface{}{
		Duck{Name: "Waffle"},
		Owl{Name: "Mr.Who"},
		Fish{Name: "Sharkbait"},
	}
	for _, anim := range animals {
		describeAnimals(anim)
	}
}
