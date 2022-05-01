/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func twoTreeSymmetric(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	return twoTreeSymmetric(a.Left, b.Right) && twoTreeSymmetric(a.Right, b.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return twoTreeSymmetric(root.Left, root.Right)
}