func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Cell [3]int

type PriorityQueue struct {
	nums []Cell
}

func Constructor() *PriorityQueue {
	return &PriorityQueue{nums: make([]Cell, 0)}
}

func (pq *PriorityQueue) up() {
	p := len(pq.nums) - 1
	for p > 0 && pq.nums[(p-1)/2][2] > pq.nums[p][2] {
		pq.nums[(p-1)/2], pq.nums[p] = pq.nums[p], pq.nums[(p-1)/2]
		p = (p - 1) / 2
	}
}

func (pq *PriorityQueue) down() {
	p := 0
	for {
		minidx := p
		if p*2+1 < len(pq.nums) && pq.nums[p*2+1][2] < pq.nums[minidx][2] {
			minidx = p*2 + 1
		}
		if p*2+2 < len(pq.nums) && pq.nums[p*2+2][2] < pq.nums[minidx][2] {
			minidx = p*2 + 2
		}
		if minidx != p {
			pq.nums[p], pq.nums[minidx] = pq.nums[minidx], pq.nums[p]
			p = minidx
		} else {
			break
		}
	}
}

func (pq *PriorityQueue) Push(cell [3]int) {
	pq.nums = append(pq.nums, cell)
	pq.up()
}

func (pq *PriorityQueue) Pop() Cell {
	n := len(pq.nums) - 1
	pq.nums[0], pq.nums[n] = pq.nums[n], pq.nums[0]
	popv := pq.nums[n]
	pq.nums = pq.nums[:n]
	pq.down()
	return popv
}

func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	n := len(heightMap[0])
	rains := make([][]int, m)
	for i := 0; i < m; i++ {
		rains[i] = make([]int, n)
		for j := 0; j < n; j++ {
			rains[i][j] = -1
		}
	}
	pq := Constructor()
	for i := 0; i < m; i++ {
		rains[i][0] = heightMap[i][0]
		rains[i][n-1] = heightMap[i][n-1]
		pq.Push(Cell([3]int{i, 0, rains[i][0]}))
		pq.Push(Cell([3]int{i, n - 1, rains[i][n-1]}))
	}
	for j := 1; j < n-1; j++ {
		rains[0][j] = heightMap[0][j]
		rains[m-1][j] = heightMap[m-1][j]
		pq.Push(Cell([3]int{0, j, rains[0][j]}))
		pq.Push(Cell([3]int{m - 1, j, rains[m-1][j]}))
	}

	for len(pq.nums) > 0 {
		cell := pq.Pop()
		if cell[0] > 0 && rains[cell[0]-1][cell[1]] == -1 {
			rains[cell[0]-1][cell[1]] = max(cell[2], heightMap[cell[0]-1][cell[1]])
			pq.Push(Cell([3]int{cell[0] - 1, cell[1], rains[cell[0]-1][cell[1]]}))
		}
		if cell[1] > 0 && rains[cell[0]][cell[1]-1] == -1 {
			rains[cell[0]][cell[1]-1] = max(cell[2], heightMap[cell[0]][cell[1]-1])
			pq.Push(Cell([3]int{cell[0], cell[1] - 1, rains[cell[0]][cell[1]-1]}))
		}
		if cell[0] < m-1 && rains[cell[0]+1][cell[1]] == -1 {
			rains[cell[0]+1][cell[1]] = max(cell[2], heightMap[cell[0]+1][cell[1]])
			pq.Push(Cell([3]int{cell[0] + 1, cell[1], rains[cell[0]+1][cell[1]]}))
		}
		if cell[1] < n-1 && rains[cell[0]][cell[1]+1] == -1 {
			rains[cell[0]][cell[1]+1] = max(cell[2], heightMap[cell[0]][cell[1]+1])
			pq.Push(Cell([3]int{cell[0], cell[1] + 1, rains[cell[0]][cell[1]+1]}))
		}
	}
	rain := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rain += rains[i][j] - heightMap[i][j]
		}
	}
	return rain
}