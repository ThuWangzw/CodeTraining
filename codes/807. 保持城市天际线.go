func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	rowheight := make([]int, n)
	colheight := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rowheight[i] = max(rowheight[i], grid[i][j])
			colheight[j] = max(colheight[j], grid[i][j])
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans += min(rowheight[i], colheight[j]) - grid[i][j]
		}
	}
	return ans
}