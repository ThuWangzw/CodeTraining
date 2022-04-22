type Node struct {
	children []string
	weight   []float64
	group    int
	value    float64
}

func newNode() *Node {
	return &Node{
		children: make([]string, 0),
		weight:   make([]float64, 0),
		group:    -1,
		value:    -1,
	}
}

func bfs(nodeid string, graph map[string]*Node, value float64, group int) {
	graph[nodeid].value = value
	graph[nodeid].group = group
	for i, childid := range graph[nodeid].children {
		if graph[childid].group == -1 {
			bfs(childid, graph, value/graph[nodeid].weight[i], group)
		}
	}
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]*Node)
	// construct graph
	for i, equation := range equations {
		s, t := equation[0], equation[1]
		if _, ok := graph[s]; !ok {
			graph[s] = newNode()
		}
		if _, ok := graph[t]; !ok {
			graph[t] = newNode()
		}
		graph[s].children = append(graph[s].children, t)
		graph[t].children = append(graph[t].children, s)
		graph[s].weight = append(graph[s].weight, values[i])
		graph[t].weight = append(graph[t].weight, 1/values[i])
	}

	// bfs
	group := 0
	for nodeid, node := range graph {
		if node.group == -1 {
			bfs(nodeid, graph, 1, group)
			group++
		}
	}

	// query
	queryn := len(queries)
	ans := make([]float64, 0, queryn)
	for _, query := range queries {
		s, t := query[0], query[1]
		if graph[s] == nil || graph[t] == nil || graph[s].group != graph[t].group {
			ans = append(ans, -1)
		} else {
			ans = append(ans, graph[s].value/graph[t].value)
		}
	}
	return ans
}