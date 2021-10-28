func isSymIter(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return isSymIter(a.Left, b.Right) && isSymIter(a.Right, b.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymIter(root.Left, root.Right)
}