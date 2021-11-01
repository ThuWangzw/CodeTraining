func preorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	queue := make([]*TreeNode, 0)
	p := root
	for p != nil || len(queue) > 0 {
		if p != nil {
			queue = append(queue, p)
			ans = append(ans, p.Val)
			p = p.Left
		} else {
			p = queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			p = p.Right
		}
	}
	return ans
}