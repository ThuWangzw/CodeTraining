func recoverTree(root *TreeNode) {
	var lastVisit *TreeNode
	var nodei *TreeNode
	var nodej *TreeNode
	p := root
	queue := make([]*TreeNode, 0)
	for p != nil || len(queue) > 0 {
		for p != nil {
			queue = append(queue, p)
			p = p.Left
		}
		p = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if lastVisit != nil {
			if p.Val < lastVisit.Val {
				if nodei == nil {
					nodei = lastVisit
					nodej = p
				} else {
					nodej = p
				}
			}
		}
		lastVisit = p
		p = p.Right
	}
	nodei.Val, nodej.Val = nodej.Val, nodei.Val
}