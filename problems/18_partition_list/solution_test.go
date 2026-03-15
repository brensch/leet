package partition_list

import (
	"reflect"
	"testing"
)

func listFromSlice(values []int) *ListNode {
	var head *ListNode
	var tail *ListNode

	for _, value := range values {
		node := &ListNode{Val: value}
		if head == nil {
			head = node
			tail = node
			continue
		}
		tail.Next = node
		tail = node
	}

	return head
}

func sliceFromList(head *ListNode) []int {
	var out []int
	for node := head; node != nil; node = node.Next {
		out = append(out, node.Val)
	}
	return out
}

func TestPartition(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		x    int
		want []int
	}{
		{
			name: "example case",
			in:   []int{1, 4, 3, 2, 5, 2},
			x:    3,
			want: []int{1, 2, 2, 4, 3, 5},
		},
		{
			name: "single pivot move",
			in:   []int{2, 1},
			x:    2,
			want: []int{1, 2},
		},
		{
			name: "all less than x",
			in:   []int{1, 1, 1},
			x:    5,
			want: []int{1, 1, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := sliceFromList(Partition(listFromSlice(tc.in), tc.x))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Partition(%v, %d) = %v, want %v", tc.in, tc.x, got, tc.want)
			}
		})
	}
}
