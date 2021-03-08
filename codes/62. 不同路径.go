// 方法一，用组合数来算
func uniquePaths(m int, n int) int {
	large := m - 1 + n - 1
	small := int(math.Max(float64(m-1), float64(n-1)))
	var result uint64 = 1
	var tmp uint64 = 1
	j := 1
	for i := large; i > small; i-- {
		result *= uint64(i)
		tmp *= uint64(j)
		j++
	}
	return int(result / tmp)
}

// 方法二：用动态规划
func uniquePaths(m int, n int) int {
	dp_path := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		dp_path = append(dp_path, make([]int, n, n))
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp_path[i][j] = 1
			} else {
				dp_path[i][j] = dp_path[i-1][j] + dp_path[i][j-1]
			}
		}
	}
	return dp_path[m-1][n-1]
}