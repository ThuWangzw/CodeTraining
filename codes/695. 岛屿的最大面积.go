// BFS
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func BFS(grid [][]int, i int, j int) int {
	m := len(grid)
	n := len(grid[0])
	area := 1
	grid[i][j] = 0
	if i > 0 && grid[i-1][j] == 1 {
		area += BFS(grid, i-1, j)
	}
	if i < m-1 && grid[i+1][j] == 1 {
		area += BFS(grid, i+1, j)
	}
	if j > 0 && grid[i][j-1] == 1 {
		area += BFS(grid, i, j-1)
	}
	if j < n-1 && grid[i][j+1] == 1 {
		area += BFS(grid, i, j+1)
	}
	return area
}

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	maxArea := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, BFS(grid, i, j))
			}
		}
	}
	return maxArea
}