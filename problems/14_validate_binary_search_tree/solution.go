package validate_binary_search_tree

import "math"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// IsValidBST returns whether the tree satisfies BST rules.
func IsValidBST(root *TreeNode) bool {

	return validInRange(root, math.MinInt, math.MaxInt)

}

func validInRange(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}

	leftValid := validInRange(root.Left, min, root.Val)
	rightValid := validInRange(root.Right, root.Val, max)

	if !leftValid || !rightValid {
		return false
	}
	return true

}
