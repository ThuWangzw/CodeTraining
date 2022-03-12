func postorder(root *Node) []int {
	if root == nil {
		return make([]int, 0)
	}
	stack := []*Node{root}
	visited := make(map[*Node]bool)
	ans := make([]int, 0)
	for len(stack) > 0 {
		n := len(stack)
		top := stack[n-1]
		if len(top.Children) == 0 || visited[top.Children[len(top.Children)-1]] {
			visited[top] = true
			stack = stack[:n-1]
			ans = append(ans, top.Val)
		} else {
			for i := len(top.Children) - 1; i >= 0; i-- {
				stack = append(stack, top.Children[i])
			}
		}
	}
	return ans
}