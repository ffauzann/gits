package main

import (
	"fmt"
	"os"
	"regexp"
)

/*
	example valid input:
		matched:
			{([][]())[{}]}
			{({[][]})}
			{}{}{()}
			[[()]]
		unmatched:
			[[[[[({})(()}]]]]]
			{([][]{())}[{}]}
			{(}
	otherwise:
	{a}
	a
*/

func main() {
	var bracket string
	fmt.Print("Input bracket(s)	: ")
	fmt.Scan(&bracket)

	if valid := ValidateInput(bracket); !valid {
		fmt.Println("\nInvalid input")
		os.Exit(0)
	}

	balanced, reason := BracketPair(bracket)
	fmt.Println("Is balanced		:", balanced)
	fmt.Println("Reason			:", reason)

}

// BracketPair returns whether given bracket is balanced
func BracketPair(s string) (balanced bool, reason string) {
	if len(s)%2 == 1 || len(s) == 0 {
		reason = "given bracket(s) is odd or empty"
		return
	}

	half := len(s) / 2
	for i := range s {
		opposite, l, r, j := OppositeBracket(s[i]), i-1, i+1, i+3
		if (i > 0 && s[l] == opposite) || (r < len(s) && s[r] == opposite) {
			continue
		}

		if i > half {
			j = i - 3
		}

		if (i == 0 && opposite != s[r]) || (i > 0 && opposite != s[l]) {
			fmt.Println("not matched yet", i, string(s[i]), string(opposite))
			for {
				if j >= len(s) || j < 0 {
					reason = fmt.Sprintf("given bracket on position %d makes unbalance", i+1)
					return
				}

				if opposite == s[j] {
					fmt.Println("found matched at", i, j, string(s[i]), string(opposite))
					break
				}

				if i > half {
					j -= 2
					continue
				}
				j += 2
			}
		}
	}

	reason = "All bracket(s) balanced perfectly"
	return true, reason
}

// ValidateInput return whether the input is valid
func ValidateInput(s string) bool {
	return !regexp.MustCompile(`[^\{\}\[\]\(\)]`).MatchString(s)
}

// OppositeBracket returns byte of the opposite of given bracket
func OppositeBracket(b byte) (opposite byte) {
	switch string(b) {
	case "{":
		opposite = byte('}')
	case "}":
		opposite = byte('{')
	case "(":
		opposite = byte(')')
	case ")":
		opposite = byte('(')
	case "[":
		opposite = byte(']')
	case "]":
		opposite = byte('[')
	}

	return
}
