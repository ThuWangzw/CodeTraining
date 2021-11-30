func increasingBSTIter(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	var lstart, lend, rstart, rend *TreeNode
	if root.Left != nil {
		lstart, lend = increasingBSTIter(root.Left)
	}
	if root.Right != nil {
		rstart, rend = increasingBSTIter(root.Right)
	}
	start := lstart
	end := lend
	if lend != nil {
		end.Left = nil
		end.Right = root
		end = root
		start.Left = nil
		end.Left = nil
	} else {
		start = root
		root.Left = nil
		end = root
	}
	if rstart == nil {
		return start, end
	} else {
		rstart.Left = nil
		end.Right = rstart
		end = rend
		end.Right = nil
		return start, end
	}
}

func increasingBST(root *TreeNode) *TreeNode {
	root, _ = increasingBSTIter(root)
	return root
}