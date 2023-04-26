package pathsum

import (
	"reflect"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	tests := map[string]struct {
		root     *TreeNode
		sum      int
		expected bool
	}{
		`first`: {
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 2}},
			sum:      9,
			expected: true,
		},
		`second`: {
			root:     &TreeNode{Val: 1},
			sum:      1,
			expected: true,
		},
		`third`: {
			root:     &TreeNode{Val: 2},
			sum:      1,
			expected: false,
		},
		`forth`: {
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 2}},
			sum:      4,
			expected: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := hasPathSum(tt.root, tt.sum)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got: %v, expected: %v", got, tt.expected)
			}
		})
	}
}
