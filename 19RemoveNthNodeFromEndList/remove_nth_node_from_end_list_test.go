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
		inp := tt.input
		n := tt.n
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := removeNthFromEnd(inp, n)
			if !equals(got, exp) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
