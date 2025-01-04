package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Age int
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("User is underage: %d", v.Age)
}

func validateAge(age int) (bool, error) {
	if age < 0 {
		return false, errors.New("Age is negative, try again.")
	}
	if age >= 18 {
		return true, nil
	}

	return false, &ValidationError{Age: age}
}

func main() {
	fmt.Println("Testing age over 18")
	ret, err := validateAge(100)
	fmt.Println("Value of result", ret, "and value of error", err)
	fmt.Println("Testing age under 18")
	ret, err = validateAge(1)
	fmt.Println("Value of result", ret, "and value of error", err)
}
