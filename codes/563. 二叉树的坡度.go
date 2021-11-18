func findTiltIter(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftSum, leftTilt := findTiltIter(root.Left)
	rightSum, rightTiil := findTiltIter(root.Right)
	return leftSum + rightSum + root.Val, int(math.Abs(float64(leftSum-rightSum))) + leftTilt + rightTiil
}

func findTilt(root *TreeNode) int {
	_, tilt := findTiltIter(root)
	return tilt
}