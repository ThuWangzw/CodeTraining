func countPairs(n int, edges [][]int) int64 {
	graph := make([][]int, n)
	flag := make([]bool, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	cnt := 0
	var ans int64
	for i := 0; i < n; i++ {
		if !flag[i] {
			cur := bfs(graph, i, flag)
			ans += int64(cur) * int64(cnt)
			cnt += cur
		}
	}
	return ans
}

func bfs(graph [][]int, node int, flag []bool) int {
	queue := make([]int, 0)
	queue = append(queue, node)
	flag[node] = true
	cnt := 0
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		cnt++
		for _, child := range graph[top] {
			if !flag[child] {
				queue = append(queue, child)
				flag[child] = true
			}
		}
	}
	return cnt
}