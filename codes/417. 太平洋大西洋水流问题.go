func BFS(start []int, heights [][]int, traveled [][]bool, flag [][]int) {
	if traveled[start[0]][start[1]] {
		return
	}
	m := len(heights)
	n := len(heights[0])
	traveled[start[0]][start[1]] = true
	queue := make([][]int, 1)
	queue[0] = start
	for len(queue)>0 {
		cur := queue[0]
		queue = queue[1:]
		flag[cur[0]][cur[1]] += 1
		if cur[0]>0 && (!traveled[cur[0]-1][cur[1]]) && heights[cur[0]-1][cur[1]]>=heights[cur[0]][cur[1]] {
			traveled[cur[0]-1][cur[1]] = true
			queue = append(queue, []int{cur[0]-1, cur[1]})
		}
		if cur[0]<m-1 && (!traveled[cur[0]+1][cur[1]]) && heights[cur[0]+1][cur[1]]>=heights[cur[0]][cur[1]] {
			traveled[cur[0]+1][cur[1]] = true
			queue = append(queue, []int{cur[0]+1, cur[1]})
		}
		if cur[1]>0 && (!traveled[cur[0]][cur[1]-1]) && heights[cur[0]][cur[1]-1]>=heights[cur[0]][cur[1]] {
			traveled[cur[0]][cur[1]-1] = true
			queue = append(queue, []int{cur[0], cur[1]-1})
		}
		if cur[1]<n-1 && (!traveled[cur[0]][cur[1]+1]) && heights[cur[0]][cur[1]+1]>=heights[cur[0]][cur[1]] {
			traveled[cur[0]][cur[1]+1] = true
			queue = append(queue, []int{cur[0], cur[1]+1})
		}
	}
}

func pacificAtlantic(heights [][]int) [][]int {
	m := len(heights)
	n := len(heights[0])
	traveled := make([][]bool, m)
	for i:=0; i<m; i++ {
		traveled[i] = make([]bool, n)
	}
	flag := make([][]int, m)
	for i:=0; i<m; i++ {
		flag[i] = make([]int, n)
	}
	results := make([][]int, 0)
	for i:=0; i<n; i++ {
		BFS([]int{0,i}, heights, traveled, flag,)
	}
	for i:=0; i<m; i++ {
		BFS([]int{i,0}, heights, traveled, flag)
	}
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++{
			traveled[i][j] = false
		}
	}
	for i:=0; i<n; i++ {
		BFS([]int{m-1,i}, heights, traveled, flag)
	}
	for i:=0; i<m; i++ {
		BFS([]int{i,n-1}, heights, traveled, flag)
	}
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if flag[i][j]==2 {
				results = append(results, []int{i, j})
			}
		}
	}
	return results
}