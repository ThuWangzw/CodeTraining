/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isDirectSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	return A.Val == B.Val && isDirectSubStructure(A.Left, B.Left) && isDirectSubStructure(A.Right, B.Right)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if A.Val == B.Val {
		flag := isDirectSubStructure(A, B)
		if flag {
			return true
		}
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}