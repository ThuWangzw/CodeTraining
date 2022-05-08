type Node []int

func xtosum(x int) int {
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return sum
}

func movingCount(m int, n int, k int) int {
	flag := make([][]bool, m)
	for i := 0; i < m; i++ {
		flag[i] = make([]bool, n)
	}
	queue := make([]Node, 0)
	flag[0][0] = true
	queue = append(queue, []int{0, 0})
	cnt := 1
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		directions := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
		for _, direction := range directions {
			node := []int{direction[0] + top[0], direction[1] + top[1]}
			if node[0] >= 0 && node[0] < m && node[1] >= 0 && node[1] < n && !flag[node[0]][node[1]] && xtosum(node[0])+xtosum(node[1]) <= k {
				queue = append(queue, node)
				flag[node[0]][node[1]] = true
				cnt++
			}
		}
	}
	return cnt
}