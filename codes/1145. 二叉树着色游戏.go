/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var flag bool = false

func checkValid(a, b, c int) bool {
	if b >= a && b >= c {
		a, b = b, a
	}
	if c >= a && c >= b {
		a, c = c, a
	}
	return a > b+c+1
}

func calculate(root *TreeNode, n int, x int) (int, bool) {
	if root == nil {
		return 0, false
	}
	lc, end := calculate(root.Left, n, x)
	if end {
		return 0, true
	}
	rc, end := calculate(root.Right, n, x)
	if end {
		return 0, true
	}
	if root.Val == x {
		flag = checkValid(n-1-lc-rc, lc, rc)
		return 0, true
	}
	return lc + rc + 1, false
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	calculate(root, n, x)
	return flag
}