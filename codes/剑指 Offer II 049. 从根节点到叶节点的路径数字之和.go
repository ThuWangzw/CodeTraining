/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var sum int = 0

func sumNumbersIter(root *TreeNode, cur int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		sum += cur*10 + root.Val
		return
	}
	sumNumbersIter(root.Left, cur*10+root.Val)
	sumNumbersIter(root.Right, cur*10+root.Val)
}

func sumNumbers(root *TreeNode) int {
	sum = 0
	sumNumbersIter(root, 0)
	return sum
}