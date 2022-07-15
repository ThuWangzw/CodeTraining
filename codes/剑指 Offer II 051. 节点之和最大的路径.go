/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var ans int

func maxPathSum(root *TreeNode) int {
	ans = math.MinInt32
	maxBranchSum(root)
	return ans
}

func maxBranchSum(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	lml, lmr := maxBranchSum(root.Left)
	rml, rmr := maxBranchSum(root.Right)
	ml := max(max(lml, lmr), 0)
	mr := max(max(rml, rmr), 0)
	ml += root.Val
	mr += root.Val
	ans = max(ans, mr+ml-root.Val)
	return ml, mr
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}