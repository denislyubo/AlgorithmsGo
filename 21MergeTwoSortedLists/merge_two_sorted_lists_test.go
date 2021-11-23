package mergetwosortedlists

import "testing"

func isEqual(l1 *ListNode, l2 *ListNode) bool {
	for {
		if l1 == nil && l2 == nil {
			return true
		} else if (l1 == nil && l2 != nil) || (l2 == nil && l1 != nil) {
			return false
		}

		if l1.Val != l2.Val {
			return false
		}

		l1 = l1.Next
		l2 = l2.Next
	}
}

func TestMergeTwoSortedLists(t *testing.T) {
	l1 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  7,
			Next: nil,
		},
	}

	l2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	l3 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
		},
	}

	l4 := mergeTwoLists(l1, l2)

	if !isEqual(l3, l4) {
		t.Error("Error")
	}

}
