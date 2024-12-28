package main

import "fmt"

func main() {
	name := "Genesis"
	numCats := 4
	wantsMoreCats := true

	formattedString := fmt.Sprintf("My name is %s and I have %d cats. It is %t, I want more cats.", name, numCats, wantsMoreCats)
	fmt.Println(formattedString)
}
