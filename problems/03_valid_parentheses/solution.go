package valid_parentheses

import "fmt"

// IsValid reports whether every opening bracket is closed in the correct order.
func IsValid(s string) bool {
	// each thing needs its corresponding close.

	// make arrays with the open first and closed second

	// 0th element = opener, 1st = closer
	pairings := [][]rune{
		{'(', ')'},
		{'[', ']'},
		{'{', '}'},
	}
	fmt.Println("--------------------")
	fmt.Println(s)

	runeStack := []rune{}
	for _, currentRune := range s {
		for _, pairing := range pairings {

			// if empty stack and opener, add
			if len(runeStack) == 0 {
				if currentRune == pairing[0] {
					runeStack = append(runeStack, currentRune)
					break
				}
				// if empty and trying to close, bad
				if currentRune == pairing[1] {
					return false
				}
			}

			// if last thing on the stack was the opener for this closer, remove opener
			if len(runeStack) > 0 && runeStack[len(runeStack)-1] == pairing[0] &&
				currentRune == pairing[1] {
				runeStack = runeStack[0 : len(runeStack)-1]
				break
			}

			// if last thing on the stack was not the opener for this closer and this is an opener, add to stack
			if currentRune == pairing[0] {
				runeStack = append(runeStack, currentRune)
				break
			}

			// if this is a closer at this point then it was wrong
			if currentRune == pairing[1] {
				return false
			}
		}
	}

	if len(runeStack) != 0 {
		return false
	}
	return true
}
