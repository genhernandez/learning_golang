package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeToFile(userInput string) error {
	path := "./user_input.txt"
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(userInput)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter text to write to file:")
	scanner.Scan()
	text := scanner.Text()
	fmt.Println("Saved:", text)
	err := writeToFile(text)
	if err != nil {
		err = fmt.Errorf("failed to write '%s' to user_input.txt", text)
		fmt.Print(err.Error())
	} else {
		fmt.Println("Successfully saved message", text, "to user_input.txt")
	}

}
