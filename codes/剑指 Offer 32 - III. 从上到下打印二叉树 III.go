/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	ans := make([][]int, 0)
	queue := []*TreeNode{root}
	reverse := false
	for len(queue) > 0 {
		nextQueue := make([]*TreeNode, 0)
		row := make([]int, 0)
		for i := 0; i < len(queue); i++ {
			idx := i
			if reverse {
				idx = len(queue) - i - 1
			}
			row = append(row, queue[idx].Val)
			if queue[i].Left != nil {
				nextQueue = append(nextQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				nextQueue = append(nextQueue, queue[i].Right)
			}
		}
		queue = nextQueue
		ans = append(ans, row)
		reverse = !reverse
	}
	return ans
}