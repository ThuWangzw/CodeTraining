func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[n-1]}
	leftlen := 0
	for inorder[leftlen] != root.Val {
		leftlen++
	}
	root.Left = buildTree(inorder[:leftlen], postorder[:leftlen])
	root.Right = buildTree(inorder[leftlen+1:], postorder[leftlen:n-1])
	return root
}