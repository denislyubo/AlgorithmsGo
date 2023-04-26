package removenthnodefromendlist

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	tests := map[string]struct {
		input    *ListNode
		n        int
		expected *ListNode
	}{
		`first`: {
			input: &ListNode{1,
				&ListNode{2,
					&ListNode{3,
						nil},
				}},
			n: 1,
			expected: &ListNode{1,
				&ListNode{2,
					nil},
			},
		},
		`second`: {
			input: &ListNode{1,
				&ListNode{2,
					&ListNode{3,
						&ListNode{4,
							nil},
					}}},
			n: 3,
			expected: &ListNode{1,
				&ListNode{3,
					&ListNode{4,
						nil},
				}},
		},
		`third`: {
			input: &ListNode{1,
				nil},
			n:        1,
			expected: nil,
		},
		`fourth`: {
			input: &ListNode{1,
				&ListNode{2,
					nil},
			},
			n: 1,
			expected: &ListNode{1,
				nil},
		},
		`fiftieth`: {
			input: &ListNode{1,
				&ListNode{2,
					&ListNode{3,
						&ListNode{4,
							&ListNode{5,
								nil}},
					}}},
			n: 1,
			expected: &ListNode{1,
				&ListNode{2,
					&ListNode{3,
						&ListNode{4,
							nil}},
				}},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			inp := tt.input
			n := tt.n
			got := removeNthFromEnd(inp, n)
			exp := tt.expected
			if !equals(got, exp) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
