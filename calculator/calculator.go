package main

import (
	"flag"
	"fmt"
)

func addition(a, b int64) {
	total := a + b
	fmt.Printf("Total of sum: %d\n", total)
}

func subtraction(a, b int64) {
	total := a - b
	fmt.Printf("Total of subtraction: %d\n", total)
}

func multiplication(a, b int64) {
	total := a * b
	fmt.Printf("Total of multiplication: %d\n", total)
}

func division(a, b int64) {
	if b == 0 {
		fmt.Println("Error: Division by zero is not allowed")
		return
	}

	total := float64(a) / float64(b)
	fmt.Printf("Total of division: %.2f\n", total)
}

func main() {
	var num1 int64
	var num2 int64

	flag.Int64Var(&num1, "n1", 0, "First Number")
	flag.Int64Var(&num2, "n2", 0, "Second Number")

	flag.Parse()

	addition(num1, num2)
	subtraction(num1, num2)
	multiplication(num1, num2)
	division(num1, num2)
}
