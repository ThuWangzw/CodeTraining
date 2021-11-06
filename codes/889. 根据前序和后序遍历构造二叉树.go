func constructTreeIter(preorder []int, postorder []int) (bool, *TreeNode) {
	if len(preorder) == 0 {
		return true, nil
	}
	n := len(preorder)
	if len(preorder) != len(postorder) || preorder[0] != postorder[n-1] {
		return false, nil
	}
	root := &TreeNode{Val: preorder[0]}
	for i := 0; i <= n-1; i++ {
		ok, leftChild := constructTreeIter(preorder[1:1+i], postorder[:i])
		if !ok {
			continue
		}
		ok, rightChild := constructTreeIter(preorder[i+1:n], postorder[i:n-1])
		if ok {
			root.Left = leftChild
			root.Right = rightChild
			return true, root
		}
	}
	return false, nil
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	_, root := constructTreeIter(preorder, postorder)
	return root
}