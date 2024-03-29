package reverselinkedlist

import "testing"

func TestReverseLinkedList(t *testing.T) {
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
			expected: &ListNode{3,
				&ListNode{2,
					&ListNode{1,
						nil},
				}},
		},
	}

	for name, tt := range tests {
		inp := tt.input
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := reverseList(inp)
			if !equals(got, exp) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
