func minPathCost(grid [][]int, moveCost [][]int) int {
	m := len(grid)
	n := len(grid[0])
	cost := make([][]int, m)
	for i := 0; i < m; i++ {
		cost[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		cost[0][i] = grid[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			cost[i][j] = math.MaxInt32
			for k := 0; k < n; k++ {
				cost[i][j] = min(cost[i][j], moveCost[grid[i-1][k]][j]+cost[i-1][k]+grid[i][j])
			}
		}
	}
	mincost := math.MaxInt32
	for i := 0; i < n; i++ {
		mincost = min(mincost, cost[m-1][i])
	}
	return mincost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}