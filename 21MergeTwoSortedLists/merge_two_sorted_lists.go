package mergetwosortedlists

/**
 * Definition for singly-linked list.
 */

// ListNode is storage struct for list
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil && l2 != nil {
		return l2
	} else if l2 == nil && l1 != nil {
		return l1
	}

	var head *ListNode

	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}

	var prev *ListNode = head

	for {
		if l1 == nil {
			prev.Next = l2
			break
		}
		if l2 == nil {
			prev.Next = l1
			break
		}

		if l1.Val < l2.Val {
			prev.Next = l1
			prev = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			prev = l2
			l2 = l2.Next
		}
	}

	return head
}
