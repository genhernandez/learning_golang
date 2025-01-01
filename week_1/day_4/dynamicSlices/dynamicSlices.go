package main

import "fmt"

func filterByThreshold(int_slice []int, threshold int) []int {
	max_capacity := cap(int_slice)
	fmt.Println("Creating result slice with max capacity", max_capacity)
	result := make([]int, 0, max_capacity)

	for _, v := range int_slice {
		if v <= threshold {
			fmt.Println(v, "is less than or equal to ", threshold)
			result = append(result, v)
		}
	}
	return result
}

func main() {
	original_slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, -1, -3, -5}
	fmt.Println("Original slice", original_slice)
	threshold := 9

	final_slice := filterByThreshold(original_slice, threshold)
	fmt.Println("Final slice", final_slice)
}
