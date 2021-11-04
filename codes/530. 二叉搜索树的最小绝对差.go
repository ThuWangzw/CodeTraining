func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func getMinimumDifference(root *TreeNode) int {
	last := -1
	mindiff := math.MaxInt32
	p := root
	stack := make([]*TreeNode, 0)
	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if last == -1 {
			last = p.Val
		} else {
			mindiff = min(p.Val-last, mindiff)
			last = p.Val
		}
		p = p.Right
	}
	return mindiff
}