/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func visit(root *TreeNode, ans *int) (int, int) {
	if root == nil {
		return 0, 0
	}
	sum, cnt := root.Val, 1
	if root.Left != nil {
		lsum, lcnt := visit(root.Left, ans)
		sum += lsum
		cnt += lcnt
	}
	if root.Right != nil {
		rsum, rcnt := visit(root.Right, ans)
		sum += rsum
		cnt += rcnt
	}
	if sum/cnt == root.Val {
		(*ans)++
	}
	return sum, cnt
}

func averageOfSubtree(root *TreeNode) int {
	ans := 0
	visit(root, &ans)
	return ans
}