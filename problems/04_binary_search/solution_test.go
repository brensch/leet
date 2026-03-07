package binary_search

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "found", nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, want: 4},
		{name: "missing", nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, want: -1},
		{name: "single element", nums: []int{5}, target: 5, want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Search(tc.nums, tc.target)
			if got != tc.want {
				t.Fatalf("Search(%v, %d) = %d, want %d", tc.nums, tc.target, got, tc.want)
			}
		})
	}
}
