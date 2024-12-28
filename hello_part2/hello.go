package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	helloString := "Hello, World!"
	nameString := "My name is Genesis"
	dateString := fmt.Sprintf("Current time:  %s", currentTime.String())
	fmt.Println(helloString)
	fmt.Println(nameString)
	fmt.Println(dateString)
}
