package mergetwobinarytree

// TreeNode is binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	node := &TreeNode{}

	if root1 != nil {
		node.Val += root1.Val
	}
	if root2 != nil {
		node.Val += root2.Val
	}
	var l1, l2, r1, r2 *TreeNode
	if root1 != nil {
		l1 = root1.Left
		r1 = root1.Right
	}
	if root2 != nil {
		l2 = root2.Left
		r2 = root2.Right
	}
	node.Left = mergeTrees(l1, l2)
	node.Right = mergeTrees(r1, r2)

	return node
}

func equal(root1 *TreeNode, root2 *TreeNode) bool {
	if (root1 == nil && root2 == nil) ||
		(root1 != nil && root2 != nil && root1.Val == root2.Val &&
			equal(root1.Left, root2.Left) &&
			equal(root1.Right, root2.Right)) {
		return true
	}
	return false
}
