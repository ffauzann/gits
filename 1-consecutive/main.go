package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Print("Input number (separated by comma with no space): ")
	fmt.Scan(&input)
	consecutive := ConsecutiveNumber(ParseInput(input))
	fmt.Println(consecutive)
}

// ParseInput returns slice of int from the given input
func ParseInput(s string) (ints []int) {
	split := strings.Split(s, ",")
	for idx := range split {
		i, err := strconv.Atoi(split[idx])
		if err != nil {
			fmt.Println("Invalid input:", split[idx])
			os.Exit(0)
		}

		if i != 0 && i != 1 {
			fmt.Println("Invalid input: only 0 and 1 are allowed")
			os.Exit(0)
		}

		ints = append(ints, i)
	}

	return
}

// ConsecutiveNumber returns the number of consecutive by given slice of number
func ConsecutiveNumber(numbers []int) (maxConsecutive int) {
	var currentConsecutive int
	for i := range numbers {
		if i == 0 {
			maxConsecutive++
			currentConsecutive++
			continue
		}

		if numbers[i] == numbers[i-1] {
			currentConsecutive++
			if currentConsecutive > maxConsecutive {
				maxConsecutive = currentConsecutive
			}
			continue
		}

		currentConsecutive = 1
	}

	return
}
