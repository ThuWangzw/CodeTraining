func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	step := (min(m, n) + 1) / 2
	ans := make([]int, 0, m*n)
	for k := 0; k < step; k++ {
		for j := k; j < n-k; j++ {
			ans = append(ans, matrix[k][j])
		}
		for i := k + 1; i < m-k; i++ {
			ans = append(ans, matrix[i][n-k-1])
		}
		if m-k-1 != k && n-k-1 != k {
			for j := n - k - 2; j >= k; j-- {
				ans = append(ans, matrix[m-k-1][j])
			}
			for i := m - k - 2; i >= k+1; i-- {
				ans = append(ans, matrix[i][k])
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}