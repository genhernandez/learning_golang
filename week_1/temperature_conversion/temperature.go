package main

import (
	"flag"
	"fmt"
)

func main() {
	var fahrenheitDegrees float64
	flag.Float64Var(&fahrenheitDegrees, "f", 0.0, "Temperature in Fahrenheit")

	flag.Parse()

	celsius := (5.0 / 9.0) * (fahrenheitDegrees - 32)

	fmt.Printf("Fahrenheit: %.2f\n", fahrenheitDegrees)
	fmt.Printf("Celsius: %.2f\n", celsius)
}
