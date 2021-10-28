func delNodesIter(root *TreeNode, to_delete map[int]bool) (*TreeNode, []*TreeNode) {
	if root == nil {
		return nil, make([]*TreeNode, 0)
	}
	if to_delete[root.Val] {
		lchild, ltrees := delNodesIter(root.Left, to_delete)
		rchild, rtrees := delNodesIter(root.Right, to_delete)
		trees := append(ltrees, rtrees...)
		if lchild != nil {
			trees = append(trees, lchild)
		}
		if rchild != nil {
			trees = append(trees, rchild)
		}
		return nil, trees
	} else {
		lchild, ltrees := delNodesIter(root.Left, to_delete)
		rchild, rtrees := delNodesIter(root.Right, to_delete)
		root.Left = lchild
		root.Right = rchild
		return root, append(ltrees, rtrees...)
	}
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	deleteMap := make(map[int]bool)
	for _, num := range to_delete {
		deleteMap[num] = true
	}
	root, trees := delNodesIter(root, deleteMap)
	if root != nil {
		trees = append(trees, root)
	}
	return trees
}