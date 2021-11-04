type Item struct {
	node  *TreeNode
	depth int
}

func findBottomLeftValue(root *TreeNode) int {
	bottomVal := -1
	bottomDepth := -1
	p := &Item{node: root, depth: 0}
	stack := make([]*Item, 0)
	for p != nil || len(stack) > 0 {
		for p != nil {
			if p.depth > bottomDepth {
				bottomDepth = p.depth
				bottomVal = p.node.Val
			}
			stack = append(stack, p)
			if p.node.Left != nil {
				p = &Item{node: p.node.Left, depth: p.depth + 1}
			} else {
				p = nil
			}
		}
		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if p.node.Right != nil {
			p = &Item{node: p.node.Right, depth: p.depth + 1}
		} else {
			p = nil
		}
	}
	return bottomVal
}