package maximum_depth_of_binary_tree

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "single node",
			root: &TreeNode{Val: 42},
			want: 1,
		},
		{
			name: "balanced tree",
			root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			want: 3,
		},
		{
			name: "skewed tree",
			root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
			want: 3,
		},
		{
			name: "left skewed tree",
			root: &TreeNode{Val: 8, Left: &TreeNode{Val: 8, Left: &TreeNode{Val: 8, Left: &TreeNode{Val: 8}}}},
			want: 4,
		},
		{
			name: "uneven branches",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{Val: 6},
					},
				},
			},
			want: 4,
		},
		{
			name: "same values different shape",
			root: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 7,
					Right: &TreeNode{Val: 7},
				},
			},
			want: 3,
		},
		{name: "empty tree", root: nil, want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MaxDepth(tc.root)
			if got != tc.want {
				t.Fatalf("MaxDepth(...) = %d, want %d", got, tc.want)
			}
		})
	}
}
