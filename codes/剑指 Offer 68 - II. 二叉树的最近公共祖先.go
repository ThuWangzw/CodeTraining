/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestorIter(root, p, q *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}
	rootflag := 0
	if root.Val == p.Val || root.Val == q.Val {
		rootflag = 1
	}
	leftflag, leftAns := lowestCommonAncestorIter(root.Left, p, q)
	if leftAns != nil {
		return 1, leftAns
	}
	rightflag, rightAns := lowestCommonAncestorIter(root.Right, p, q)
	if rightAns != nil {
		return 1, rightAns
	}
	flag := rootflag + leftflag + rightflag
	if flag == 2 {
		return 1, root
	} else if flag == 1 {
		return 1, nil
	} else {
		return 0, nil
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, ans := lowestCommonAncestorIter(root, p, q)
	return ans
}