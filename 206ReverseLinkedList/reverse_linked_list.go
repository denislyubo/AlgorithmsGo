package reverselinkedlist

// ListNode is definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev, nxt *ListNode
	curr := head

	for curr != nil {
		nxt = curr.Next
		curr.Next = prev
		prev = curr
		curr = nxt
	}

	return prev
}

func equals(left, right *ListNode) bool {
	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	if left != nil || right != nil {
		return false
	}

	return true
}
