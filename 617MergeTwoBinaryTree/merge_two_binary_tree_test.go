package mergetwobinarytree

import (
	"testing"
)

func TestMergeTrees(t *testing.T) {
	tests := map[string]struct {
		left     *TreeNode
		right    *TreeNode
		expected *TreeNode
	}{
		`first`: {
			left:     &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 2}},
			right:    &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}},
			expected: &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 7}}},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := mergeTrees(tt.left, tt.right)
			exp := tt.expected
			if !equal(tt.expected, got) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
