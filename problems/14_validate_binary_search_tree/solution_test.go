package validate_binary_search_tree

import "testing"

func TestIsValidBST(t *testing.T) {

	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "valid bst",
			root: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: true,
		},
		{
			name: "invalid because descendant breaks rule",
			root: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 4},
				Right: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: false,
		},
		{
			name: "invalid because left subtree has too-large descendant",
			root: &TreeNode{
				Val:  10,
				Left: &TreeNode{Val: 5, Right: &TreeNode{Val: 12}},
				Right: &TreeNode{
					Val: 15,
				},
			},
			want: false,
		},
		{
			name: "duplicate value is invalid",
			root: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: false,
		},
		{
			name: "empty tree",
			root: nil,
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsValidBST(tc.root)
			if got != tc.want {
				t.Fatalf("IsValidBST(...) = %t, want %t", got, tc.want)
			}
		})
	}
}
