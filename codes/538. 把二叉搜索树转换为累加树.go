// 利用属性：二叉平衡树的中序遍历是有序的
func convertBSTIter(root *TreeNode, acm *int) {
	if root == nil {
		return
	}
	convertBSTIter(root.Right, acm)
	*acm += root.Val
	root.Val = *acm
	convertBSTIter(root.Left, acm)
}

func convertBST(root *TreeNode) *TreeNode {
	acm := 0
	convertBSTIter(root, &acm)
	return root
}