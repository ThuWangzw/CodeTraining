func longestCycle(edges []int) int {
	n := len(edges)
	dis := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = -1
	}
	ans := -1
	for i := 0; i < n; i++ {
		ans = max(ans, findCycle(edges, i, dis))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findCycle(edges []int, node int, dis []int) int {
	visited := make(map[int]bool)
	d := 0
	for node != -1 && !visited[node] && dis[node] == -1 {
		dis[node] = d
		d++
		visited[node] = true
		node = edges[node]
	}
	if visited[node] {
		return d - dis[node]
	}
	return -1
}