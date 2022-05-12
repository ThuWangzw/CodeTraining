/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBalancedIter(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	lb, ld := isBalancedIter(root.Left)
	if !lb {
		return false, 0
	}
	rb, rd := isBalancedIter(root.Right)
	if !rb {
		return false, 0
	}
	return math.Abs(float64(ld-rd)) < 2, max(ld, rd) + 1
}

func isBalanced(root *TreeNode) bool {
	ok, _ := isBalancedIter(root)
	return ok
}