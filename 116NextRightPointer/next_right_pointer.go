package nextrightpointer

import (
	"algogo/Queue"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var q queue.GenQueue[*Node]
	q.PushBack(root)
	for !q.Empty() {
		n := q.Size()
		a := n

		for i := 0; i < a; i++ {
			temp := q.Top()
			q.PopFront()
			if i == a-1 {
				temp.Next = nil
			} else {
				temp.Next = q.Top()
			}

			if temp.Left != nil {
				q.PushBack(temp.Left)
			}

			if temp.Right != nil {
				q.PushBack(temp.Right)
			}
		}
	}

	return root
}

func equal(root1 *Node, root2 *Node) bool {
	if (root1 == nil && root2 == nil) ||
		(root1 != nil && root2 != nil && root1.Val == root2.Val &&
			equal(root1.Left, root2.Left) &&
			equal(root1.Right, root2.Right) &&
			equal(root1.Next, root2.Next)) {
		return true
	}
	return false
}
