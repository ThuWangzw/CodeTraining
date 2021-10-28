func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func depthAndDiameter(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	ldepth, lmax := depthAndDiameter(root.Left)
	rdepth, rmax := depthAndDiameter(root.Right)
	depth := max(ldepth, rdepth) + 1
	diameter := ldepth + rdepth
	diameter = max(diameter, max(lmax, rmax))
	return depth, diameter
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, dia := depthAndDiameter(root)
	return dia
}