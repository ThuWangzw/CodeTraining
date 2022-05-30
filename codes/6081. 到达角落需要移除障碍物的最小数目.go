func minimumObstacles(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	distance := make([][]int, m)
	for i := 0; i < m; i++ {
		distance[i] = make([]int, n)
		for j := 0; j < n; j++ {
			distance[i][j] = -1
		}
	}
	begindis := 0
	if grid[0][0] == 1 {
		begindis = 1
	}
	distance[0][0] = begindis
	begins := [][]int{{0, 0}}
	for len(begins) > 0 {
		nexts := make([][]int, 0)
		for _, begin := range begins {
			travel(begin, grid, distance, begindis, &nexts)
		}
		begindis++
		begins = nexts
	}
	return distance[m-1][n-1]
}

// 从一个0开始遍历，为所有相同距离的0开始赋值，并将相邻的1赋值为distance+1，修改为0，压入队列,返回一个相邻的1
func travel(begin []int, grid [][]int, distance [][]int, depth int, neighs *[][]int) {
	m := len(grid)
	n := len(grid[0])
	directions := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	for _, direction := range directions {
		next := []int{begin[0] + direction[0], begin[1] + direction[1]}
		if next[0] >= 0 && next[0] < m && next[1] >= 0 && next[1] < n && distance[next[0]][next[1]] == -1 {
			if grid[next[0]][next[1]] == 0 {
				distance[next[0]][next[1]] = depth
				travel(next, grid, distance, depth, neighs)
			} else {
				distance[next[0]][next[1]] = depth + 1
				grid[next[0]][next[1]] = 0
				*neighs = append(*neighs, []int{next[0], next[1]})
			}
		}
	}
}