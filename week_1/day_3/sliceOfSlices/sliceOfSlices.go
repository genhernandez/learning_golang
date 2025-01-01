package main

import "fmt"

func printSubSlices(nestedSlice [][]int) {
	for i := 0; i < len(nestedSlice); i++ {
		fmt.Println("Printing slice", i)
		for j := 0; j < len(nestedSlice[i]); j++ {
			fmt.Println("Element", j, "in slice", i, "is", nestedSlice[i][j])
		}
	}
}

func main() {
	nestedSlice := [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	fmt.Println("Length of nested slice is", len(nestedSlice))

	printSubSlices(nestedSlice)
}
