func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	left := make([][]int, m)
	for i := 0; i < m; i++ {
		left[i] = make([]int, n)
	}
	leftCnt := 0
	for i := 0; i < m; i++ {
		leftCnt = 0
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				leftCnt = 0
			} else {
				leftCnt++
				left[i][j] = leftCnt
			}
		}
	}
	maxarea := 0
	for i := 0; i < n; i++ {
		monoStack := make([]int, 0)
		areas := make([]int, m)
		for j := 0; j < m; j++ {
			leftbound := -1
			for len(monoStack) > 0 && left[monoStack[len(monoStack)-1]][i] >= left[j][i] {
				monoStack = monoStack[:len(monoStack)-1]
			}
			if len(monoStack) > 0 {
				leftbound = monoStack[len(monoStack)-1]
			}
			monoStack = append(monoStack, j)
			areas[j] = (j - leftbound) * left[j][i]
		}
		monoStack = make([]int, 0)
		for j := m - 1; j >= 0; j-- {
			rightbound := m
			for len(monoStack) > 0 && left[monoStack[len(monoStack)-1]][i] >= left[j][i] {
				monoStack = monoStack[:len(monoStack)-1]
			}
			if len(monoStack) > 0 {
				rightbound = monoStack[len(monoStack)-1]
			}
			monoStack = append(monoStack, j)
			areas[j] += (rightbound - j - 1) * left[j][i]
		}
		for j := 0; j < m; j++ {
			maxarea = max(maxarea, areas[j])
		}
	}
	return maxarea
}