package removenthnodefromendlist

// ListNode is definition of slightly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{Next: head}
	fast, slow := temp, temp

	for hit := 1; hit <= n; hit++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return temp.Next
}

func equals(l *ListNode, r *ListNode) bool {
	if l == nil && r == nil {
		return true
	}

	for ; l != nil && r != nil; l, r = l.Next, r.Next {
		if (l == nil && r != nil) ||
			(r == nil && l != nil) ||
			(l.Val != r.Val) {
			return false
		}
	}

	return true
}
