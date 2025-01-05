package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(srcPath string) (string, error) {
	fmt.Println("Opening", srcPath)
	f, err := os.Open(srcPath)

	if err != nil {
		err = fmt.Errorf("error opening %s: %w", srcPath, err)
		return "", err
	}
	defer f.Close()
	content, err := io.ReadAll(f)
	if err != nil {
		err = fmt.Errorf("error reading %s: %w", srcPath, err)
		return "", err
	}
	fmt.Println("Successfully read from", srcPath)
	return string(content), nil
}

func writeFile(dstPath, message string) error {
	f, err := os.OpenFile(dstPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		err = fmt.Errorf("error opening %s: %w", dstPath, err)
		return err
	}
	defer f.Close()
	_, writeErr := f.WriteString(message)
	if writeErr != nil {
		writeErr = fmt.Errorf("error writing to %s: %w", dstPath, writeErr)
		return writeErr
	}
	fmt.Println("Successfully wrote to", dstPath)
	return nil
}

func copyFileWithBuffer(srcPath, dstPath string, bufferSize int) error {
	srcFile, err := os.Open(srcPath)

	if err != nil {
		return fmt.Errorf("error opening source file %s: %w", srcPath, err)
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dstPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening destination file %s: %w", dstPath, err)
	}
	defer dstFile.Close()

	buf := make([]byte, bufferSize)

	for {
		n, readErr := srcFile.Read(buf)
		if n > 0 {
			if _, writeErr := dstFile.Write(buf[:n]); writeErr != nil {
				return fmt.Errorf("error writing to destination file %s: %w", dstPath, writeErr)
			}
		}
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			return fmt.Errorf("error reading from src file %s: %w", srcPath, readErr)
		}
	}
	fmt.Println("Successfully copied from", srcPath, "to", dstPath)
	return nil
}

func copyFile(srcPath, dstPath string) error {
	srcContent, err := readFile(srcPath)

	if err != nil {
		return err
	}

	writeErr := writeFile(dstPath, srcContent)

	if writeErr != nil {
		return writeErr
	}
	fmt.Println("Successfully copied from", srcPath, "to", dstPath)
	return nil
}

func main() {
	srcPath := "./file_a.txt"
	dstPath := "./file_b.txt"
	bufferSize := 4096
	if err := copyFileWithBuffer(srcPath, dstPath, bufferSize); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if err := copyFile(srcPath, dstPath); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
