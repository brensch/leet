package longest_consecutive_sequence

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "standard example",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "duplicates included",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		{
			name: "empty input",
			nums: nil,
			want: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LongestConsecutive(tc.nums)
			if got != tc.want {
				t.Fatalf("LongestConsecutive(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}
