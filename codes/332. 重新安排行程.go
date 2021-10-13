// 求欧拉图一笔画问题，倒序入栈，因为倒序入栈的顺序就是不断发现死胡同节点的顺序
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

func visit(graph map[string][]string, node string, remain int, ans *[]string) bool {

	for len(graph[node]) > 0 {
		child := graph[node][0]
		graph[node] = graph[node][1:]
		visit(graph, child, remain-1, ans)
	}
	*ans = append([]string{node}, *ans...)
	return true
}

func findItinerary(tickets [][]string) []string {
	graph := make(map[string][]string)
	inOut := make(map[string]([]int))

	for _, ticket := range tickets {
		_, ok := graph[ticket[0]]
		if !ok {
			graph[ticket[0]] = make([]string, 0)
		}
		graph[ticket[0]] = append(graph[ticket[0]], ticket[1])
		_, ok = inOut[ticket[0]]
		if !ok {
			inOut[ticket[0]] = []int{0, 0}
		}
		_, ok = inOut[ticket[1]]
		if !ok {
			inOut[ticket[1]] = []int{0, 0}
		}
		inOut[ticket[0]][1]++
		inOut[ticket[1]][0]++
	}

	for _, cities := range graph {
		sort.Slice(cities, func(i, j int) bool {
			a := cities[i]
			b := cities[j]
			an := len(a)
			bn := len(b)
			for k := 0; k < min(an, bn); k++ {
				if a[k] < b[k] {
					return true
				} else if a[k] > b[k] {
					return false
				}
			}
			return an < bn
		})
	}

	node := "JFK"
	ans := make([]string, 0)
	visit(graph, node, len(ans), &ans)
	return ans
}