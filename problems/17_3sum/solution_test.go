package three_sum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "standard example",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "no solution",
			nums: []int{1, 2, -2, -1},
			want: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ThreeSum(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("ThreeSum(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}
