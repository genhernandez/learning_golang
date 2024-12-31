package main

import "fmt"

func removeAtIndex(s []string, idx int) (s2 []string) {
	if idx < 0 || idx >= len(s) {
		fmt.Println("Index out of bounds")
		return s
	}
	fmt.Println("Element at idx", idx, "exists:", s[idx])
	return append(s[:idx], s[idx+1:]...)
}

func main() {
	slice1 := []string{"Red", "Green", "Blue"}
	fmt.Printf("Original slice is %s\n", slice1)
	fmt.Println("Adding to slice1")
	slice1 = append(slice1, "Orange")
	slice1 = append(slice1, "Yellow")
	fmt.Printf("Updated slice is %s\n", slice1)
	fmt.Println("Removing element at index 2")
	s2 := removeAtIndex(slice1, 2)
	fmt.Printf("Final slice is %s\n", s2)
}
