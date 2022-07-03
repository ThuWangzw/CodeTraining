/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	mat := make([][]int, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]int, n)
	}
	for i := 0; i <= (m-1)/2; i++ {
		for j := i; j <= n-1-i; j++ {
			mat[i][j] = -1
			if head != nil {
				mat[i][j] = head.Val
				head = head.Next
			}
		}
		for j := i + 1; j <= m-2-i && n-1-i >= i; j++ {
			mat[j][n-1-i] = -1
			if head != nil {
				mat[j][n-1-i] = head.Val
				head = head.Next
			}
		}
		for j := n - 1 - i; j >= i && i != m-1-i; j-- {
			mat[m-1-i][j] = -1
			if head != nil {
				mat[m-1-i][j] = head.Val
				head = head.Next
			}
		}
		for j := m - 1 - i - 1; j >= i+1 && n-1-i > i; j-- {
			mat[j][i] = -1
			if head != nil {
				mat[j][i] = head.Val
				head = head.Next
			}
		}
	}
	return mat
}