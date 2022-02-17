// 看上去是bfs，但是稍微分析一下时间复杂度就知道bfs的时间复杂度是指数级的，其实是一道动态规划
func knightProbability(n int, k int, row int, column int) float64 {
	p := make([][][]float64, k+1)
	directions := [][]int{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}}
	for i := 0; i <= k; i++ {
		p[i] = make([][]float64, n)
		for j := 0; j < n; j++ {
			p[i][j] = make([]float64, n)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			p[0][i][j] = 1
		}
	}
	for i := 1; i <= k; i++ {
		for j := 0; j < n; j++ {
			for q := 0; q < n; q++ {
				for _, direction := range directions {
					x := j + direction[0]
					y := q + direction[1]
					if x >= 0 && y >= 0 && x < n && y < n {
						p[i][j][q] += 0.125 * p[i-1][x][y]
					}
				}
			}
		}
	}
	return p[k][row][column]
}