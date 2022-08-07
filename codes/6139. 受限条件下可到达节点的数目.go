func reachableNodes(n int, edges [][]int, restricted []int) int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	for _, edge := range edges {
		s, e := edge[0], edge[1]
		graph[s] = append(graph[s], e)
		graph[e] = append(graph[e], s)
	}
	flags := make(map[int]bool)
	for _, num := range restricted {
		flags[num] = true
	}
	queue := make([]int, 0)
	queue = append(queue, 0)
	flags[0] = true
	cnt := 1
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		for _, child := range graph[top] {
			if !flags[child] {
				cnt++
				flags[child] = true
				queue = append(queue, child)
			}
		}
	}
	return cnt
}