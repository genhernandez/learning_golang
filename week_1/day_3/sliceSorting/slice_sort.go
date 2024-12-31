package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := []string{"Apple", "Grape", "Tomato", "Watermelon", "Banana", "Kiwi", "Zucchini"}
	fmt.Println("Original slice: ", s1)

	// with anonymous function
	// sort.Slice(s1, func(i, j int) bool {
	// 	return s1[i] > s1[j]
	// })

	// without anonymous function
	sort.Sort(sort.Reverse(sort.StringSlice(s1)))

	fmt.Println("Reverse sorted slice: ", s1)

}
