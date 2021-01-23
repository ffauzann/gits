package main

import "fmt"

func main() {
	example1 := []int{1, 1, 0, 1, 1, 1}
	example2 := []int{1, 0, 0, 1, 0, 1, 1}
	example3 := []int{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1}

	consecutive1 := ConsecutiveNumber(example1)
	consecutive2 := ConsecutiveNumber(example2)
	consecutive3 := ConsecutiveNumber(example3)

	fmt.Println(consecutive1)
	fmt.Println(consecutive2)
	fmt.Println(consecutive3)
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
