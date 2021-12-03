func leadsTodestinationIter(source int, graph [][]int, destination int, state []int) bool {
	if state[source] == 1 {
		return false
	} else if state[source] == 2 {
		return true
	}

	if len(graph[source]) == 0 {
		return source == destination
	}
	state[source] = 1
	for _, child := range graph[source] {
		if !leadsTodestinationIter(child, graph, destination, state) {
			return false
		}
	}
	state[source] = 2
	return true
}

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}
	state := make([]int, n)
	return leadsTodestinationIter(source, graph, destination, state)
}