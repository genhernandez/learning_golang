package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sliceSum(s1 []int) (sum int) {
	for _, v := range s1 {
		sum += v
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := []int{}

	fmt.Println("Enter numbers one by one. Type -1 to stop:")

	for {
		fmt.Print("Enter a number:")

		scanner.Scan()
		input := scanner.Text()

		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number")
			continue
		}

		if num == -1 {
			fmt.Println("Stopping intput as -1 was entered")
			break
		}

		numbers = append(numbers, num)
	}
	total := sliceSum(numbers)
	fmt.Println("Collected numbers: ", numbers)
	fmt.Println("Total", total)
}
