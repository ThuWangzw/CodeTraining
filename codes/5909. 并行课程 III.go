func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func computeTime(graph [][]int, target int, alltimes []int, times []int) int {
	if alltimes[target-1] >= 0 {
		return alltimes[target-1]
	}
	needTime := times[target-1]
	childMaxTime := 0
	for _, child := range graph[target] {
		childMaxTime = max(childMaxTime, computeTime(graph, child, alltimes, times))
	}
	needTime += childMaxTime
	alltimes[target-1] = needTime
	return needTime
}

func minimumTime(n int, relations [][]int, time []int) int {
	graph := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		graph[i] = make([]int, 0)
	}
	for _, relation := range relations {
		graph[relation[1]] = append(graph[relation[1]], relation[0])
	}
	alltimes := make([]int, n)
	for i := 0; i < n; i++ {
		alltimes[i] = -1
	}

	for i := 1; i <= n; i++ {
		computeTime(graph, i, alltimes, time)
	}
	maxtime := 0
	for _, mtime := range alltimes {
		maxtime = max(maxtime, mtime)
	}
	return maxtime
}