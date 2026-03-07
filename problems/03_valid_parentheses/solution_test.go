package valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "simple valid", s: "()[]{}", want: true},
		{name: "crossed nesting", s: "([)]", want: false},
		{name: "nested valid", s: "{[]}", want: true},
		{name: "missing close", s: "(", want: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsValid(tc.s)
			if got != tc.want {
				t.Fatalf("IsValid(%q) = %v, want %v", tc.s, got, tc.want)
			}
		})
	}
}
