package reverse_linked_list

import "fmt"

// ListNode is a singly linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList reverses a singly linked list.
func ReverseList(head *ListNode) *ListNode {

	var prev *ListNode
	current := head
	for {
		if current == nil {
			return prev
		}
		fmt.Println(current.Val)
		next := current.Next
		current.Next = prev
		prev = current
		current = next

	}
}
