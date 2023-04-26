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
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := middleNode(tt.input)
			if (got == nil && got != tt.expected) || (got != nil && got.Val != tt.expected.Val) {
				t.Errorf("got: %v, expected: %v", got, tt.expected)
			}
		})
	}
}
