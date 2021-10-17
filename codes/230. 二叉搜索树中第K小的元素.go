// 中序遍历
func kthSmallest(root *TreeNode, k int) int {
	p := root
	stack := make([]*TreeNode, 0)
	idx := 0
	for p != nil || len(stack) > 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			idx++
			if idx == k {
				return p.Val
			}
			p = p.Right
		}
	}
	return -1
}