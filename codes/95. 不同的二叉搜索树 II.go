// 递归
func generateTreesNodes(nodes []int) []*TreeNode {
	if len(nodes) == 0 {
		trees := make([]*TreeNode, 1)
		trees[0] = nil
		return trees
	}
	trees := make([]*TreeNode, 0)
	for i, node := range nodes {
		leftTrees := generateTreesNodes(nodes[:i])
		rightTrees := generateTreesNodes(nodes[i+1:])
		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {
				trees = append(trees, &TreeNode{Val: node, Left: leftTree, Right: rightTree})
			}
		}
	}
	return trees
}

func generateTrees(n int) []*TreeNode {
	nodes := make([]int, n, n)
	for i := 1; i <= n; i++ {
		nodes[i-1] = i
	}
	trees := generateTreesNodes(nodes)
	return trees
}