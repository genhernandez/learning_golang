package main

import "fmt"

func generateFibonacci(n int) []int {
	var seq []int
	base := []int{1, 1}
	if n < 1 {
		return seq
	}
	if n == 1 {
		seq = append(seq, 1)
		return seq
	}
	if n == 2 {
		seq = append(seq, base...)
		return seq
	}

	seq = append(seq, base...)

	for i := 2; i < n; i++ {
		prev := seq[i-1]
		before_prev := seq[i-2]
		seq = append(seq, prev+before_prev)
	}

	return seq
}

func main() {
	fib_1 := generateFibonacci(1)
	fib_2 := generateFibonacci(2)
	fib_3 := generateFibonacci(3)

	fmt.Println("Result of generateFibonacci(1) is: ", fib_1)
	fmt.Println("Result of generateFibonacci(2) is: ", fib_2)
	fmt.Println("Result of generateFibonacci(3) is: ", fib_3)

	fmt.Println("Testing on a larger number...")
	fib_12 := generateFibonacci(12)
	fmt.Println("Result of generateFibonacci(12) is: ", fib_12)
}
