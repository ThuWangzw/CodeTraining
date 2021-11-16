// 用一个标志记录上一次访问的节点，以此判断根节点是否需要访问
func postorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	cur := root
	var pre *TreeNode
	stack := make([]*TreeNode, 0)
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.Right == nil || top.Right == pre {
			ans = append(ans, top.Val)
			pre = top
			cur = nil
		} else {
			stack = append(stack, top)
			cur = top.Right
		}
	}
	return ans
}