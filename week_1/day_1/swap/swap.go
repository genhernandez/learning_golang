package main

import "fmt"

func swap(word1, word2 string) (string, string) {
	return word2, word1
}

func main() {
	original_a, original_b := "Hello", "World!"
	swapped_a, swapped_b := swap(original_a, original_b)
	formatted_string := fmt.Sprintf("Original words: %s %s. Swapped words: %s %s", original_a, original_b, swapped_a, swapped_b)
	fmt.Println(formatted_string)
}
