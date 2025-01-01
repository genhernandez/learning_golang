package main

import (
	"fmt"
)

func convertTo2D(flat_slice []int, num_rows int) [][]int {
	flat_length := len(flat_slice)
	max_row_size := (flat_length + num_rows - 1) / num_rows
	nested_slice := make([][]int, num_rows)

	for i := 0; i < flat_length; i++ {
		row := i / max_row_size
		nested_slice[row] = append(nested_slice[row], flat_slice[i])
	}
	for i := range nested_slice {
		for len(nested_slice[i]) < max_row_size {
			nested_slice[i] = append(nested_slice[i], 0)
		}
	}

	return nested_slice
}

func main() {
	flat_slice := []int{1, 2, 3, 4, 5, 6}
	num_rows := 2
	fmt.Println("Original slice", flat_slice)
	nested_slice := convertTo2D(flat_slice, num_rows)
	fmt.Println("Final slice", nested_slice)
}
