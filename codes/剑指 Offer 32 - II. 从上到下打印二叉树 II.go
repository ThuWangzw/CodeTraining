/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) [][]int {
    if root==nil {
        return make([][]int, 0)
    }
    ans := make([][]int, 0)
    queue := []*TreeNode{root}
    for len(queue)>0 {
        nextQueue := make([]*TreeNode, 0)
        row := make([]int, 0)
        for i:=0; i<len(queue); i++ {
            row = append(row, queue[i].Val)
            if queue[i].Left!=nil {
                nextQueue = append(nextQueue, queue[i].Left)
            }
            if queue[i].Right!=nil {
                nextQueue = append(nextQueue, queue[i].Right)
            }
        }
        queue = nextQueue
        ans = append(ans, row)
    }
    return ans
}