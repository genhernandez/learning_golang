package main

import "fmt"

func filterEvenNumbers(s []int) []int {
	var no_evens []int

	for _, v := range s {
		if v%2 == 0 {
			fmt.Println("Found an even value: ", v)
		} else {
			no_evens = append(no_evens, v)
		}
	}
	return no_evens
}

func main() {
	fullSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original slice: ", fullSlice)
	result := filterEvenNumbers(fullSlice)
	fmt.Println("Final filtered slice: ", result)
}
