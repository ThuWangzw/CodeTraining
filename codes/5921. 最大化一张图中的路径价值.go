func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxPathIter(values []int, graph map[int][][]int, curvalue int, curtime int, curnode int, visited map[int]bool, maxvalue *int, maxTime int) {
	if curnode == 0 {
		*maxvalue = max(*maxvalue, curvalue)
	}
	for _, child := range graph[curnode] {
		if child[1]+curtime <= maxTime {
			cv := curvalue
			originvisit := visited[child[0]]
			if !visited[child[0]] {
				visited[child[0]] = true
				cv += values[child[0]]
			}
			maxPathIter(values, graph, cv, child[1]+curtime, child[0], visited, maxvalue, maxTime)
			visited[child[0]] = originvisit
		}
	}
}

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	n := len(values)
	graph := make(map[int][][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([][]int, 0)
	}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], []int{edge[1], edge[2]})
		graph[edge[1]] = append(graph[edge[1]], []int{edge[0], edge[2]})
	}
	visited := make(map[int]bool)
	visited[0] = true
	cv := values[0]
	maxv := 0
	maxPathIter(values, graph, cv, 0, 0, visited, &maxv, maxTime)
	return maxv
}