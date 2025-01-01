package main

import (
	"fmt"
)

func convertTo2D(flat_slice []int, num_rows int) [][]int {
	flat_length := len(flat_slice)
	max_row_size := (flat_length + num_rows - 1) / num_rows
	nested_slice := make([][]int, num_rows)
	fmt.Println("Creating a 2d slice with num_rows", num_rows, "and max_row_size", max_row_size)
	for i := 0; i < flat_length; i++ {
		row := i / max_row_size
		nested_slice[row] = append(nested_slice[row], flat_slice[i])
	}
	for i := range nested_slice {
		for len(nested_slice[i]) < max_row_size {
			fmt.Println("Padding with extra zeros")
			nested_slice[i] = append(nested_slice[i], 0)
		}
	}
	return nested_slice
}

func main() {
	flat_slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	num_rows := 4
	fmt.Println("Original slice", flat_slice)
	nested_slice := convertTo2D(flat_slice, num_rows)
	fmt.Println("Final slice", nested_slice)
}
