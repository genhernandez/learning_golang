package main

import (
	"fmt"
	"io"
	"os"
)

func readFiles(filePaths []string) ([]string, []error) {
	fileContents := make([]string, 0, len(filePaths))
	fileErrors := make([]error, 0, len(filePaths))

	for _, filePath := range filePaths {
		func() {
			fmt.Println("Trying to open", filePath)
			f, err := os.Open(filePath)

			if err != nil {
				fileErrors = append(fileErrors, err)
				return
			}
			defer f.Close()
			content, err := io.ReadAll(f)
			if err != nil {
				fileErrors = append(fileErrors, err)
				return
			}
			fileContents = append(fileContents, string(content))
		}()
	}
	return fileContents, fileErrors
}

func main() {
	filePaths := []string{
		"/Users/genesishernandez/Desktop/learning_golang/week_1/day_5/multipleErrors/first.txt",
		"/Users/genesishernandez/Desktop/learning_golang/week_1/day_5/multipleErrors/second.txt",
		"/Users/genesishernandez/Desktop/learning_golang/week_1/day_5/multipleErrors/third.txt",
		"/Users/genesishernandez/Desktop/learning_golang/week_1/day_5/multipleErrors/fourth.txt", // does not exist
		"/Users/genesishernandez/Desktop/learning_golang/week_1/day_5/multipleErrors/fifth.txt",  // does not exist
	}

	fileContents, fileErrors := readFiles(filePaths)

	fmt.Println("File contents:")
	for _, content := range fileContents {
		fmt.Println(content)
	}

	fmt.Println("\nErrors:")
	for _, err := range fileErrors {
		fmt.Println(err)
	}
}
