func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	min_path := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		min_path = append(min_path, make([]int, n, n))
		if i == 0 {
			for j := 0; j < n; j++ {
				if j == 0 {
					min_path[i][j] = grid[i][j]
				} else {
					min_path[i][j] = grid[i][j] + min_path[i][j-1]
				}
			}
		} else {
			for j := 0; j < n; j++ {
				if j == 0 {
					min_path[i][j] = min_path[i-1][j] + grid[i][j]
				} else {
					min_path[i][j] = int(math.Min(float64(min_path[i-1][j]), float64(min_path[i][j-1]))) + grid[i][j]
				}
			}
		}
	}
	return min_path[m-1][n-1]
}