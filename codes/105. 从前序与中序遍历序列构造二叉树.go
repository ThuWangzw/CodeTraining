func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	leftend := 0
	for ; leftend < n; leftend++ {
		if inorder[leftend] == preorder[0] {
			break
		}
	}
	leftlen := leftend
	rightlen := n - leftend - 1
	if leftlen != 0 {
		root.Left = buildTree(preorder[1:1+leftlen], inorder[:leftlen])
	}
	if rightlen != 0 {
		root.Right = buildTree(preorder[1+leftlen:], inorder[leftlen+1:])
	}
	return root
}