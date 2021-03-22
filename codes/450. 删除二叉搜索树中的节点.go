func deleteNode(root *TreeNode, key int) *TreeNode {
	if root==nil {
		return root
	}
	if root.Val==key {
		// find
		if root.Right==nil {
			return root.Left
		}
		
		new_root := root.Right
		new_left := new_root.Left
		new_root.Left = root.Left
		if new_root.Left==nil {
			new_root.Left = new_left
		} else {
			p := new_root.Left
			for p.Right!=nil {
				p = p.Right
			}
			p.Right = new_left
		}
		return new_root	
	}
	if root.Val>key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}