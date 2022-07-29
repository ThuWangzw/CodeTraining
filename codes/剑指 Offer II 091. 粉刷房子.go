func minCost(costs [][]int) int {
	n := len(costs)
	allcosts := make([][]int, n)
	allcosts[0] = []int{costs[0][0], costs[0][1], costs[0][2]}
	for i := 1; i < n; i++ {
		allcosts[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			allcosts[i][j] = math.MaxInt32
			for k := 0; k < 3; k++ {
				if j != k {
					allcosts[i][j] = min(allcosts[i][j], allcosts[i-1][k]+costs[i][j])
				}
			}
		}
	}
	return min(allcosts[n-1][0], min(allcosts[n-1][1], allcosts[n-1][2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}