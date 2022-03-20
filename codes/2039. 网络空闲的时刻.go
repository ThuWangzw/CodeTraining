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

func networkBecomesIdle(edges [][]int, patience []int) int {
	n := len(patience)
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	// bfs
	distance := make([]int, n)
	arrived := make([]bool, n)
	arrived[0] = true
	queue := []int{0}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		for _, child := range graph[top] {
			if !arrived[child] {
				distance[child] = distance[top] + 1
				arrived[child] = true
				queue = append(queue, child)
			}
		}
	}

	maxtime := 0
	for i := 1; i < n; i++ {
		dis := distance[i]
		maxtime = max(maxtime, patience[i]*((dis*2-1)/patience[i])+dis*2+1)
	}
	return maxtime
}