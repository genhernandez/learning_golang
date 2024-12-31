package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := []int{10, 3, 1, 8, 9, 2}

	fmt.Println("Original slice: ", s1)

	sort.Ints(s1) // Sorts in ascending order

	fmt.Println("Sorted slice: ", s1)
}
