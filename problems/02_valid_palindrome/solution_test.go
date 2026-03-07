package valid_palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "phrase with punctuation", s: "A man, a plan, a canal: Panama", want: true},
		{name: "not a palindrome", s: "race a car", want: false},
		{name: "only punctuation", s: "!!!", want: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsPalindrome(tc.s)
			if got != tc.want {
				t.Fatalf("IsPalindrome(%q) = %v, want %v", tc.s, got, tc.want)
			}
		})
	}
}
