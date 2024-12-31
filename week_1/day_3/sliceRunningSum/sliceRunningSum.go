package main

import "fmt"

func runningSum(s1 []int) {
	sum := 0

	for _, v := range s1 {
		sum += v
		fmt.Println("Current sum: ", sum)
	}
	fmt.Println("Final sum: ", sum)
}

func main() {
	s1 := []int{1, 2, 3, 4, 5, 12, 23, 13, 1, 4, 31, 1}
	fmt.Println("Original slice: ", s1)

	runningSum(s1)
}
