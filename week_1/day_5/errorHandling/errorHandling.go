package main

import (
	"errors"
	"fmt"
)

func divide(dividend, divisor int) (int, error) {
	quotient := 0

	if divisor == 0 {
		return quotient, errors.New("Cannot divide by 0")
	}
	fmt.Println("Dividing", dividend, "by", divisor)
	return (dividend / divisor), nil
}

func main() {
	fmt.Println("Testing our divide function with 10 and 2")
	result, err := divide(10, 2)
	fmt.Println("Value of result is", result, "and value of error is", err)

	fmt.Println("Testing divide by 0")
	result_b, err_b := divide(10, 0)
	fmt.Println("Value of result is", result_b, "and value of error is", err_b)
}
