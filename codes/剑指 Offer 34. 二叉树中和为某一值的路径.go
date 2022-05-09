/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == target {
			return [][]int{{root.Val}}
		} else {
			return [][]int{}
		}
	}
	nextSum := target - root.Val
	ans := pathSum(root.Left, nextSum)
	ans = append(ans, pathSum(root.Right, nextSum)...)
	for i := 0; i < len(ans); i++ {
		ans[i] = append([]int{root.Val}, ans[i]...)
	}
	return ans
}