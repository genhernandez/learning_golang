package main

import (
	"flag"
	"fmt"
	"math"
)

func calculateCircleArea(radius float64) {
	pi := math.Pi

	area := pi * (math.Pow(radius, 2))

	fmt.Printf("Area of circle with radius %.2f: %.2f\n", radius, area)
}

func main() {
	var radius float64

	flag.Float64Var(&radius, "r", 0.0, "Radius (float) of Circle")
	flag.Parse()

	calculateCircleArea(radius)
}
