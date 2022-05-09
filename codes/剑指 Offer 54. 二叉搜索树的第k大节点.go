/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
	cnt := 0
	return kthLargestIter(root, k, &cnt)
}

func kthLargestIter(root *TreeNode, k int, cnt *int) int {
	if root == nil {
		return -1
	}
	val := kthLargestIter(root.Right, k, cnt)
	if (*cnt) >= k {
		return val
	}
	(*cnt)++
	if (*cnt) == k {
		return root.Val
	}
	val = kthLargestIter(root.Left, k, cnt)
	if (*cnt) >= k {
		return val
	}
	return -1
}