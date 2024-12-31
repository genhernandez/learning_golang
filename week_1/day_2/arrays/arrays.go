package main

import "fmt"

func arraySum(arr [5]int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	array_example := [5]int{1, 2, 3, 4, 5}
	for i, v := range array_example {
		fmt.Printf("Original value at index %d is %d\n", i, v)
	}
	fmt.Println("Getting sum of all elements")
	sum := arraySum(array_example)
	fmt.Printf("Total of all elements: %d\n", sum)
}
