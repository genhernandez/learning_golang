package main

import "fmt"

func reverseSlice(int_slice []int) {
	end := len(int_slice) - 1
	for i := 0; i < len(int_slice)/2; i++ {
		int_slice[i], int_slice[end] = int_slice[end], int_slice[i]
		fmt.Println("slice after swap", int_slice)
		end -= 1
	}
}

func main() {
	original_slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Original slice", original_slice)
	reverseSlice(original_slice)
	fmt.Println("Reversed slice", original_slice)

}
