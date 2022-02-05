func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(grid [][]int, x int, y int) int {
	if grid[x][y] == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	directions := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	goldincell := grid[x][y]
	grid[x][y] = 0
	maxgold := goldincell
	for _, direction := range directions {
		xx := x + direction[0]
		yy := y + direction[1]
		if xx >= 0 && xx < m && yy >= 0 && yy < n {
			maxgold = max(maxgold, goldincell+dfs(grid, xx, yy))
		}
	}
	grid[x][y] = goldincell
	return maxgold
}

func getMaximumGold(grid [][]int) int {
	maxgold := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxgold = max(maxgold, dfs(grid, i, j))
		}
	}
	return maxgold
}