func sumOfLeftLeavesIter(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil && isLeft {
		return root.Val
	}
	return sumOfLeftLeavesIter(root.Left, true) + sumOfLeftLeavesIter(root.Right, false)
}

func sumOfLeftLeaves(root *TreeNode) int {
	return sumOfLeftLeavesIter(root, false)
}