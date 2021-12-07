func validCell(m, n, row, col int) bool {
	return row >= 0 && row < m && col >= 0 && col < n
}

func isBorder(grid [][]int, row int, col int) bool {
	m := len(grid)
	n := len(grid[0])
	if row == 0 || row == m-1 || col == 0 || col == n-1 {
		return true
	}
	if grid[row-1][col] != grid[row][col] ||
		grid[row+1][col] != grid[row][col] ||
		grid[row][col-1] != grid[row][col] ||
		grid[row][col+1] != grid[row][col] {
		return true
	}
	return false
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m := len(grid)
	n := len(grid[0])
	flag := make([][]bool, m)
	for i := 0; i < m; i++ {
		flag[i] = make([]bool, n)
	}
	queue := make([][]int, 0)
	candidates := make([][]int, 0)
	queue = append(queue, []int{row, col})
	flag[row][col] = true
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if isBorder(grid, front[0], front[1]) {
			candidates = append(candidates, []int{front[0], front[1]})
		}
		row = front[0]
		col = front[1]
		if validCell(m, n, row-1, col) && !flag[row-1][col] && grid[row-1][col] == grid[row][col] {
			queue = append(queue, []int{row - 1, col})
			flag[row-1][col] = true
		}
		if validCell(m, n, row+1, col) && !flag[row+1][col] && grid[row+1][col] == grid[row][col] {
			queue = append(queue, []int{row + 1, col})
			flag[row+1][col] = true
		}
		if validCell(m, n, row, col-1) && !flag[row][col-1] && grid[row][col-1] == grid[row][col] {
			queue = append(queue, []int{row, col - 1})
			flag[row][col-1] = true
		}
		if validCell(m, n, row, col+1) && !flag[row][col+1] && grid[row][col+1] == grid[row][col] {
			queue = append(queue, []int{row, col + 1})
			flag[row][col+1] = true
		}
	}
	for _, point := range candidates {
		grid[point[0]][point[1]] = color
	}
	return grid
}