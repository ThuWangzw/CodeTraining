func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node := root
	stack := make([]*TreeNode, 0)
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			if (node.Val-p.Val)*(node.Val-q.Val) <= 0 {
				return node
			}
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	return nil
}