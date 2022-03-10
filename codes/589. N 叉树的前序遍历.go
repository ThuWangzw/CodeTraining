func preorder(root *Node) []int {
	if root == nil {
		return make([]int, 0)
	}
	stack := []*Node{root}
	ans := make([]int, 0)
	for len(stack) > 0 {
		n := len(stack)
		top := stack[n-1]
		stack = stack[:n-1]
		ans = append(ans, top.Val)
		for i := len(top.Children) - 1; i >= 0; i-- {
			stack = append(stack, top.Children[i])
		}
	}
	return ans
}