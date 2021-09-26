func min(a int64, b int64) int64 {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func gridGame(grid [][]int) int64 {
	n := len(grid[0])
	sums := make([][]int64, 2)
	sums[0] = make([]int64, n)
	sums[1] = make([]int64, n)
	sums[0][0] = int64(grid[0][0])
	sums[1][0] = int64(grid[1][0])
	for i := 1; i < n; i++ {
		sums[0][i] = sums[0][i-1] + int64(grid[0][i])
		sums[1][i] = sums[1][i-1] + int64(grid[1][i])
	}

	var ans int64
	ans = math.MaxInt64
	for i := 0; i < n; i++ {
		a := sums[0][n-1] - sums[0][i]
		var b int64
		if i != 0 {
			b = sums[1][i-1]
		}
		cur := max(a, b)
		ans = min(ans, cur)
	}
	return ans
}