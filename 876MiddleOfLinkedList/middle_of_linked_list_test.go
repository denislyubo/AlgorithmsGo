package middleoflinkedlist

import (
	"testing"
)

func TestMiddleNode(t *testing.T) {
	tests := map[string]struct {
		input    *ListNode
		expected *ListNode
	}{
		`first`: {
			input: &ListNode{1,
				&ListNode{2,
					&ListNode{3,
						nil},
				}},
			expected: &ListNode{Val: 2},
		},
		`second`: {
			input: &ListNode{1,
				&ListNode{2,
					&ListNode{3,
						&ListNode{4,
							nil},
					}}},
			expected: &ListNode{Val: 3},
		},
		`third`: {
			input: &ListNode{1,
				nil},
			expected: &ListNode{Val: 1},
		},
		`forth`: {
			input: nil,
		},
	}

	for name, tt := range tests {
		inp := tt.input
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := middleNode(inp)
			if (got == nil && got != exp) || (got != nil && got.Val != exp.Val) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
