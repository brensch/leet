package maximum_depth_of_binary_tree

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MaxDepth returns the maximum depth of the binary tree.
func MaxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	leftMax := MaxDepth(root.Left)
	rightMax := MaxDepth(root.Right)
	if leftMax > rightMax {
		return leftMax + 1
	} else {
		return rightMax + 1
	}
}
