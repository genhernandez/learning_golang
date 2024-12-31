package main

import "fmt"

func main() {
	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	firstFive := s2[:5]
	lastFive := s2[5:]
	middle := s2[3:8]

	fmt.Println("Original slice: ", s2)
	fmt.Printf("Capacity of original slice: %d\n", cap(s2))
	fmt.Println("First five slice:", firstFive)
	fmt.Printf("Capacity of firstFive: %d\n", cap(firstFive))
	fmt.Println("Last five slice:", lastFive)
	fmt.Printf("Capacity of lastFive: %d\n", cap(lastFive))
	fmt.Println("Middle slice:", middle)
	fmt.Printf("Capacity of middle: %d\n", cap(middle))

	fmt.Println("Adding to middle slice")
	middle = append(middle, s2...)
	fmt.Println("Middle slice after append:", middle)
	fmt.Printf("Capacity of middle after append: %d\n", cap(middle))
	fmt.Println("Original slice: ", s2)
	fmt.Println("First five slice:", firstFive)
	fmt.Println("Last five slice:", lastFive)

	fmt.Println("Adding five to each element in the original slice")
	for idx, v := range s2 {
		s2[idx] = v + 5
	}

	fmt.Println("Original slice: ", s2)
	fmt.Println("First five slice:", firstFive)
	fmt.Println("Last five slice:", lastFive)
	fmt.Println("Middle slice:", middle)

}
