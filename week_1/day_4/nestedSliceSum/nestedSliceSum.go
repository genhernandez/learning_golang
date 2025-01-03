package main

import "fmt"

func sumNestedSlice(nested_slice [][]int) []int {
	sums := make([]int, 0, len(nested_slice))
	for i := range nested_slice {
		total_sum := 0
		for _, el := range nested_slice[i] {
			total_sum += el
		}
		sums = append(sums, total_sum)
	}
	return sums
}

func main() {
	nested_slice := [][]int{{1, 2, 3}, {4, 5}, {6}, {1, 2, 9}}
	fmt.Println("Original nested slice", nested_slice)
	flat_sums := sumNestedSlice(nested_slice)
	fmt.Println("1d slice of sums", flat_sums)

}
