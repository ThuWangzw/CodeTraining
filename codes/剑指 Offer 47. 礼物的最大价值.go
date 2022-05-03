func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	vals := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		vals = append(vals, make([]int, n))
	}
	vals[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		vals[i][0] = vals[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		vals[0][i] = vals[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			vals[i][j] = max(vals[i-1][j], vals[i][j-1]) + grid[i][j]
		}
	}
	return vals[m-1][n-1]
}