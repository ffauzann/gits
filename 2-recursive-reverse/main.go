package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input	: ")
	s, _ := reader.ReadString('\n')

	reversed := RecursiveReverse(s, 0, "")
	fmt.Printf("Result	: %s \n", reversed)
}

func RecursiveReverse(s string, i int, tmp string) string {
	if i == len(s) {
		return tmp[1:]
	}

	return RecursiveReverse(s, i+1, tmp+string(s[len(s)-i-1]))
}
