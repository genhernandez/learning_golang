package main

import "fmt"

func main() {
	var fahrenheitDegrees float64 = 98.6
	celsius := (5.0 / 9.0) * (fahrenheitDegrees - 32)

	fmt.Printf("Fahrenheit: %.2f\n", fahrenheitDegrees)
	fmt.Printf("Celsius: %.2f\n", celsius)
}
