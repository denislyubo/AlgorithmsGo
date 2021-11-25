package invertbinarytree

import "testing"

func TestInvertBinaryTree(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	}

	expectedTree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	res := invertTree(tree)

	if res.Val != expectedTree.Val ||
		res.Left.Val != expectedTree.Left.Val ||
		res.Right.Val != expectedTree.Right.Val {
		t.Error("error: inversion is incorrect")
	}

}
