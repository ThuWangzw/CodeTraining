var mod int = 1e9 + 7

func countPathsIter(x, y, last int, grid [][]int, ans [][]int) int {
	if grid[x][y] <= last {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	if ans[x][y] >= 0 {
		return ans[x][y]
	}
	cnt := 1
	directions := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for _, direction := range directions {
		nx, ny := x+direction[0], y+direction[1]
		if nx < 0 || nx >= m || ny < 0 || ny >= n {
			continue
		}
		cnt = (cnt + countPathsIter(nx, ny, grid[x][y], grid, ans)) % mod
	}
	ans[x][y] = cnt
	return cnt
}

func countPaths(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ans[i][j] = -1
		}
	}
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ans[i][j] < 0 {
				countPathsIter(i, j, -1, grid, ans)
			}
			cnt = (cnt + ans[i][j]) % mod
		}
	}
	return cnt
}