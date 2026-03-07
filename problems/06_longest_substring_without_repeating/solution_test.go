package longest_substring_without_repeating

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "repeats in middle", s: "abcabcbb", want: 3},
		{name: "all same", s: "bbbbb", want: 1},
		{name: "window must jump", s: "pwwkew", want: 3},
		{name: "empty", s: "", want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LengthOfLongestSubstring(tc.s)
			if got != tc.want {
				t.Fatalf("LengthOfLongestSubstring(%q) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}
