func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

var robRes = make(map[*TreeNode]int)

func rob(root *TreeNode) int {
	if ans, ok := robRes[root]; ok {
		return ans
	}
	if root == nil {
		robRes[root] = 0
		return 0
	}
	norobroot := rob(root.Left) + rob(root.Right)
	robroot := root.Val
	if root.Left != nil {
		robroot += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		robroot += rob(root.Right.Left) + rob(root.Right.Right)
	}
	res := max(norobroot, robroot)
	robRes[root] = res
	return res
}