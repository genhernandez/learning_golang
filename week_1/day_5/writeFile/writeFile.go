package main

import (
	"fmt"
	"os"
)

func writeToFile(filePath, message string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(message)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	filePath := "outputs.txt"
	message := "Hello! I wrote this with Golang. :)"
	fmt.Println("Writing to", filePath)
	err := writeToFile(filePath, message)
	if err != nil {
		err := fmt.Errorf("problem writing to file: %w", err)
		fmt.Println(err.Error())

	} else {
		fmt.Println("Success writing to", filePath)
	}

}
