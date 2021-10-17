// 求次短路（BFS+剪枝）
func countTime(n int, time int, change int) int {
	all := 0
	for i := 0; i < n; i++ {
		if (all/change)%2 == 0 {
			// go
			all += time
		} else {
			// wait
			all += change - (all - all/change*change)
			all += time
		}
	}
	return all
}

func visit(graph map[int][]int, cur int, target int) int {
	queue := [][]int{{cur, 0}}
	minpath := make(map[int][]int)
	for node, _ := range graph {
		minpath[node] = []int{math.MaxInt32, math.MaxInt32}
	}
	for len(queue) > 0 {
		node := queue[0][0]
		deep := queue[0][1]
		queue = queue[1:]
		for _, child := range graph[node] {
			d := deep + 1
			if d < minpath[child][0] {
				minpath[child][0], minpath[child][1] = d, minpath[child][0]
				queue = append(queue, []int{child, d})
			} else if d > minpath[child][0] && d < minpath[child][1] {
				minpath[child][1] = d
				queue = append(queue, []int{child, d})
			}
		}
	}
	return minpath[target][1]
}

func secondMinimum(n int, edges [][]int, time int, change int) int {
	graph := make(map[int][]int)
	visited := make(map[int]int)
	for _, edge := range edges {
		visited[edge[0]] = math.MaxInt32
		visited[edge[1]] = math.MaxInt32
		if _, ok := graph[edge[0]]; !ok {
			graph[edge[0]] = make([]int, 0)
		}
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	visited[1] = 0
	secondpath := visit(graph, 1, n)
	return countTime(secondpath, time, change)
}