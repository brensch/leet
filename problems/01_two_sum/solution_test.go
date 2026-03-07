package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "basic case", nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{name: "pair later in slice", nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{name: "duplicate values", nums: []int{3, 3}, target: 6, want: []int{0, 1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TwoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("TwoSum(%v, %d) = %v, want %v", tc.nums, tc.target, got, tc.want)
			}
		})
	}
}
