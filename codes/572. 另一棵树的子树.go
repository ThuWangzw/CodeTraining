func isSubTreeRoot(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root == nil {
		return false
	} else if subRoot == nil {
		return false
	}
	return root.Val == subRoot.Val && isSubTreeRoot(root.Left, subRoot.Left) && isSubTreeRoot(root.Right, subRoot.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root == nil {
		return false
	} else if subRoot == nil {
		return false
	}
	if isSubTreeRoot(root, subRoot) {
		return true
	} else {
		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	}
}