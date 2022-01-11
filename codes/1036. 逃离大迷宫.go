func bfs(source []int, target []int, blocked map[int]map[int]bool) (inend bool, finded bool) {
	count := 0
	visited := make(map[int]map[int]bool)
	queue := [][]int{{source[0], source[1]}}
	visited[source[0]] = make(map[int]bool)
	visited[source[0]][source[1]] = true
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		count++
		if count >= 20000 {
			return false, false
		}
		if top[0] == target[0] && top[1] == target[1] {
			return false, true
		}
		directions := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
		for _, direction := range directions {
			node := []int{direction[0] + top[0], direction[1] + top[1]}
			if node[0] >= 0 && node[0] <= 999999 && node[1] >= 0 && node[1] <= 999999 {
				// blocked continue
				if _, err := blocked[node[0]]; err && blocked[node[0]][node[1]] {
					continue
				}
				// visited continue
				if _, err := visited[node[0]]; err && visited[node[0]][node[1]] {
					continue
				}
				// add in queue
				queue = append(queue, node)
				if _, err := visited[node[0]]; !err {
					visited[node[0]] = make(map[int]bool)
				}
				visited[node[0]][node[1]] = true
			}
		}
	}
	return true, false
}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	blockedMap := make(map[int](map[int]bool))
	for _, node := range blocked {
		if _, err := blockedMap[node[0]]; !err {
			blockedMap[node[0]] = make(map[int]bool)
		}
		blockedMap[node[0]][node[1]] = true
	}
	sourcein, finded := bfs(source, target, blockedMap)
	if finded {
		return true
	}
	targetin, finded := bfs(target, source, blockedMap)
	if finded {
		return true
	}
	return (!sourcein) && (!targetin)
}