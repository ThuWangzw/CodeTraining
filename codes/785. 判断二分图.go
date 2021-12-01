func checkPoint(id int, flag []int, graph [][]int, isvisit []bool, shouldin int) bool {
	if isvisit[id] && flag[id] != shouldin {
		return false
	}
	if isvisit[id] {
		return true
	}
	isvisit[id] = true
	flag[id] = shouldin
	for _, neighbor := range graph[id] {
		if !checkPoint(neighbor, flag, graph, isvisit, 1-shouldin) {
			return false
		}
	}
	return true
}

func isBipartite(graph [][]int) bool {
	n := len(graph)
	flag := make([]int, n)
	isvisit := make([]bool, n)
	for i := 0; i < n; i++ {
		if isvisit[i] {
			continue
		}
		if !checkPoint(i, flag, graph, isvisit, 0) {
			return false
		}
	}
	return true
}