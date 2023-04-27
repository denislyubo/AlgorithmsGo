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
		l := tt.left
		r := tt.right
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := mergeTrees(l, r)
			if !equal(exp, got) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
