package valid_palindrome

import (
	"fmt"
	"strings"
)

// IsPalindrome reports whether s is a valid palindrome after normalization.
func IsPalindrome(s string) bool {

	sCleaned := CleanString(s)
	for i := 0; i < len(sCleaned); i++ {
		if sCleaned[i] != sCleaned[len(sCleaned)-1-i] {
			return false
		}
	}
	return true
}

func CleanString(s string) string {
	start := 'a'
	end := 'z'
	sCleaned := []rune{}
	s = strings.ToLower(s)
	for _, char := range s {
		if char >= start && char <= end {
			sCleaned = append(sCleaned, char)
		}
	}
	fmt.Println(string(sCleaned))
	return string(sCleaned)
}
