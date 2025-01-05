package main

import (
	"fmt"
	"math"
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

func printShapeDetails(s Shape) {
	fmt.Println("Perimiter of shape is", s.Perimiter())
	fmt.Println("Area of shape is", s.Area())
}

func main() {
	circle := Circle{Radius: 4.0}
	rectangle := Rectangle{Width: 4.0, Height: 3.0}
	printShapeDetails(circle)
	printShapeDetails(rectangle)
}
