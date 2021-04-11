func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	neighbors := make([][]int, n)
	traveled := make([]bool, n)
	//init
	for i := 0; i < n; i++ {
		neighbor := make([]int, 0)
		for j := 0; j < n; j++ {
			if i != j && isConnected[i][j] == 1 {
				neighbor = append(neighbor, j)
			}
		}
		neighbors[i] = neighbor
	}
	cnt := 0
	// travel
	for start := 0; start < n; start++ {
		if traveled[start] {
			continue
		}
		cnt++
		traveled[start] = true
		queue := make([]int, 1)
		queue[0] = start
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, neigh := range neighbors[cur] {
				if !traveled[neigh] {
					traveled[neigh] = true
					queue = append(queue, neigh)
				}
			}
		}
	}
	return cnt
}