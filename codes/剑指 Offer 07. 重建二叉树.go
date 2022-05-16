/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	nodeVal := preorder[0]
	node := &TreeNode{
		Val: nodeVal,
	}
	for i, val := range inorder {
		if val == nodeVal {
			node.Left = buildTree(preorder[1:1+i], inorder[:i])
			node.Right = buildTree(preorder[i+1:], inorder[i+1:])
			break
		}
	}
	return node
}