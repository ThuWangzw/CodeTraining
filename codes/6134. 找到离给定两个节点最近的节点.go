func closestMeetingNode(edges []int, node1 int, node2 int) int {
	dis1 := findNodeDis(edges, node1)
	dis2 := findNodeDis(edges, node2)
	mindis, minid := math.MaxInt32, -1
	for i := 0; i < len(dis1); i++ {
		if dis1[i] == -1 || dis2[i] == -1 {
			continue
		}
		if max(dis1[i], dis2[i]) < mindis {
			mindis = max(dis1[i], dis2[i])
			minid = i
		}
	}
	return minid
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findNodeDis(edges []int, node int) []int {
	dis := make([]int, len(edges))
	for i := 0; i < len(edges); i++ {
		dis[i] = -1
	}
	d := 0
	for node != -1 && dis[node] == -1 {
		dis[node] = d
		d++
		node = edges[node]
	}
	return dis
}