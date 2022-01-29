func highestPeak(isWater [][]int) [][]int {
	m := len(isWater)
	n := len(isWater[0])
	peaks := make([][]int, m)
	for i := 0; i < m; i++ {
		peaks[i] = make([]int, n)
	}
	queue := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				queue = append(queue, []int{i, j, 0})
			}
		}
	}
	directions := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		for _, direction := range directions {
			node := []int{top[0] + direction[0], top[1] + direction[1]}
			depth := top[2]
			if node[0] >= 0 && node[0] < m && node[1] >= 0 && node[1] < n && peaks[node[0]][node[1]] == 0 && isWater[node[0]][node[1]] == 0 {
				peaks[node[0]][node[1]] = depth + 1
				queue = append(queue, []int{node[0], node[1], depth + 1})
			}
		}
	}
	return peaks
}