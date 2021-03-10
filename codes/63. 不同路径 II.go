func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	path_result := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		path_result = append(path_result, make([]int, n, n))
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				path_result[i][j] = 0
			} else if i+j == 0 {
				path_result[i][j] = 1
			} else if i == 0 {
				path_result[i][j] = path_result[i][j-1]
			} else if j == 0 {
				path_result[i][j] = path_result[i-1][j]
			} else {
				path_result[i][j] = path_result[i-1][j] + path_result[i][j-1]
			}
		}
	}
	return path_result[m-1][n-1]
}