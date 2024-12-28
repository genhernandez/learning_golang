package main

import (
	"flag"
	"fmt"
)

func main() {
	msg := "This is my message"
	var numLoops int64

	flag.Int64Var(&numLoops, "n", 1, "Number of times message should be printed")
	flag.Parse()

	for i := int64(0); i < numLoops; i++ {
		fmt.Println(msg)
	}
}
