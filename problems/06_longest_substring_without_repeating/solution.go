package longest_substring_without_repeating

// LengthOfLongestSubstring returns the size of the longest substring with unique characters.
func LengthOfLongestSubstring(s string) int {

	start := 0
	longest := 0

	for {
		if start >= len(s)-1 {
			return longest
		}

		seenChars := []byte{}
		for i := 0; i+start < len(s); i++ {
			currentLetter := s[start+i]
			foundRepeat := false
			for _, seenChar := range seenChars {
				if seenChar == currentLetter {
					foundRepeat = true
					break
				}
			}
			seenChars = append(seenChars, currentLetter)
			if foundRepeat {
				if i > longest {
					longest = i
				}
				start = start + i
				break
			}
		}

	}

}
