package main

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	Area() float64
	Perimiter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (circ Circle) Area() float64 {
	return math.Pi * math.Pow(circ.Radius, 2)
}

func (circ Circle) Perimiter() float64 {
	return 2 * math.Pi * circ.Radius
}

func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

func (rect Rectangle) Perimiter() float64 {
	return 2 * (rect.Width + rect.Height)
}

func printAnyValue(value interface{}) {
	fmt.Println("Value:", value, ", Type:", reflect.TypeOf(value))
	if shape, ok := value.(Shape); ok {
		fmt.Println("Area:", shape.Area())
		fmt.Println("Perimiter", shape.Perimiter())
	}
}

func main() {
	printAnyValue(42)
	printAnyValue(3.14)
	printAnyValue("Hello, Go!")
	printAnyValue(Rectangle{Height: 2.0, Width: 3.0})
	printAnyValue(Circle{Radius: 5.0})
}
