package main

import (
	"fmt"
	"os"
)

func readFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file does not exist: %s", filePath)
		}
		return "", fmt.Errorf("unable to open file: %s", filePath)
	}
	defer file.Close()

	// Read the file content
	stat, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("unable to retrieve file info: %s", filePath)
	}
	content := make([]byte, stat.Size())
	_, err = file.Read(content)
	if err != nil {
		return "", fmt.Errorf("unable to read file: %s", filePath)
	}

	return string(content), nil
}

func main() {
	fmt.Println("Reading from existing file")
	filePath := "/Users/genesishernandez/Desktop/gotest.rtf"
	content, err := readFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Contents:", content)
	}

	fmt.Println("Reading from nonexistent file")
	filePath = "i/dont/exist.txt"
	content, err = readFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Contents:", content)
	}
}
