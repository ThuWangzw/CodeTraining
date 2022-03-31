func flattenNode(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	if root.Left == nil && root.Right == nil {
		return root, root
	}
	leftbegin, leftend := flattenNode(root.Left)
	rightbegin, rightend := flattenNode(root.Right)
	root.Left = nil
	root.Right = nil
	end := root
	if leftbegin != nil {
		root.Right = leftbegin
		end = leftend
	}
	if rightbegin != nil {
		end.Right = rightbegin
		end = rightend
	}
	return root, end
}

func flatten(root *TreeNode) {
	flattenNode(root)
}