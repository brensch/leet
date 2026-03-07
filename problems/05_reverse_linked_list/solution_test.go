package reverse_linked_list

import (
	"reflect"
	"testing"
)

func listFromSlice(values []int) *ListNode {
	var head *ListNode
	for i := len(values) - 1; i >= 0; i-- {
		head = &ListNode{Val: values[i], Next: head}
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

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{name: "multiple nodes", in: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
		{name: "single node", in: []int{1}, want: []int{1}},
		{name: "empty list", in: nil, want: nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := sliceFromList(ReverseList(listFromSlice(tc.in)))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("ReverseList(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
