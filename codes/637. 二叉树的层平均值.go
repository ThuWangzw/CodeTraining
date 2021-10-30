type TreeNodeDepth struct {
	node  *TreeNode
	depth int
}

func averageOfLevels(root *TreeNode) []float64 {
	ans := make([]float64, 0)
	curdepth := 0
	var curval float64
	curcnt := 0
	queue := make([]*TreeNodeDepth, 0)
	queue = append(queue, &TreeNodeDepth{node: root, depth: 0})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.depth != curdepth {
			curdepth++
			ans = append(ans, curval/float64(curcnt))
			curcnt = 0
			curval = 0
		}
		curcnt++
		curval += float64(node.node.Val)
		if node.node.Left != nil {
			queue = append(queue, &TreeNodeDepth{node: node.node.Left, depth: curdepth + 1})
		}
		if node.node.Right != nil {
			queue = append(queue, &TreeNodeDepth{node: node.node.Right, depth: curdepth + 1})
		}
	}
	ans = append(ans, curval/float64(curcnt))
	return ans
}