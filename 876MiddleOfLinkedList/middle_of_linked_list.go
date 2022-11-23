package middleoflinkedlist

// ListNode is node of linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	curNode, medNode := head, head

	for curNode != nil && curNode.Next != nil {
		medNode = medNode.Next
		curNode = curNode.Next.Next
	}

	return medNode
}
