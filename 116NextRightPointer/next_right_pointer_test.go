package nextrightpointer

import (
	"testing"
)

func TestMergeTrees(t *testing.T) {
	right := &Node{Val: 3}
	r3r := &Node{Val: 7}
	r3l := &Node{Val: 6, Next: r3r}

	r2 := &Node{Val: 3,
		Left:  r3l,
		Right: r3r,
	}
	l3r := &Node{Val: 5, Next: r3l}
	exp := &Node{Val: 1,
		Left: &Node{Val: 2,
			Left:  &Node{Val: 4, Next: l3r},
			Right: l3r,
			Next:  r2},
		Right: r2,
	}
	tests := map[string]struct {
		given    *Node
		expected *Node
	}{
		`first`: {
			given:    &Node{Val: 1, Left: &Node{Val: 2}, Right: &Node{Val: 3}},
			expected: &Node{Val: 1, Left: &Node{Val: 2, Next: right}, Right: right},
		},
		`second`: {
			given: &Node{Val: 1,
				Left: &Node{Val: 2,
					Left:  &Node{Val: 4},
					Right: &Node{Val: 5}},
				Right: &Node{Val: 3,
					Left:  &Node{Val: 6},
					Right: &Node{Val: 7}}},
			expected: exp,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := connect(tt.given)
			if !equal(tt.expected, got) {
				t.Errorf("got: %v, expected: %v", got, tt.expected)
			}
		})
	}
}
