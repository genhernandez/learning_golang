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

func findLargestArea(shapes []Shape) Shape {
	if len(shapes) == 0 {
		return nil
	}
	largestAreaIdx := -1
	largestArea := math.Inf(-1)

	for idx, sh := range shapes {
		currArea := sh.Area()
		if currArea > largestArea {
			largestArea = currArea
			largestAreaIdx = idx
		}
	}
	return shapes[largestAreaIdx]
}
func printShapeDetails(s Shape) {
	fmt.Println("Perimiter of shape is", s.Perimiter())
	fmt.Println("Area of shape is", s.Area())
}

func main() {
	shapes := []Shape{
		Circle{Radius: 3.0},
		Rectangle{Width: 2.0, Height: 1.0},
		Circle{Radius: 5.0},
		Rectangle{Width: 4.0, Height: 3.0},
	}
	largestShape := findLargestArea(shapes)
	printShapeDetails(largestShape)
}
