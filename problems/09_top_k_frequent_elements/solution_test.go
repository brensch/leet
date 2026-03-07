package top_k_frequent_elements

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{name: "basic case", nums: []int{1, 1, 1, 2, 2, 3}, k: 2, want: []int{1, 2}},
		{name: "single answer", nums: []int{1}, k: 1, want: []int{1}},
		{name: "negative numbers", nums: []int{4, -1, -1, 2, -1, 2, 3}, k: 2, want: []int{-1, 2}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TopKFrequent(tc.nums, tc.k)
			sort.Ints(got)
			sort.Ints(tc.want)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("TopKFrequent(%v, %d) = %v, want %v", tc.nums, tc.k, got, tc.want)
			}
		})
	}
}
