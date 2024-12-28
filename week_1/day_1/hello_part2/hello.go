package main

import (
	"fmt"
	"time"
)

const LayoutYYYYMMDDHHMMSS = "2006-01-02 15:04:05"

func main() {
	currentTime := time.Now()
	helloString := "Hello, World!"
	nameString := "My name is Genesis"
	dateString := fmt.Sprintf("Current time:  %s", currentTime.Format(LayoutYYYYMMDDHHMMSS))
	fmt.Println(helloString)
	fmt.Println(nameString)
	fmt.Println(dateString)
}
