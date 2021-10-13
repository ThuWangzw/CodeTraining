type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])
	sums := make([][]int, m)
	for i := 0; i < m; i++ {
		sums[i] = make([]int, n)
		linesum := 0
		for j := 0; j < n; j++ {
			linesum += matrix[i][j]
			if i == 0 {
				sums[i][j] = linesum
			} else {
				sums[i][j] = sums[i-1][j] + linesum
			}
		}
	}
	return NumMatrix{sums: sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 {
		return this.sums[row2][col2]
	} else if row1 == 0 {
		return this.sums[row2][col2] - this.sums[row2][col1-1]
	} else if col1 == 0 {
		return this.sums[row2][col2] - this.sums[row1-1][col2]
	} else {
		return this.sums[row2][col2] + this.sums[row1-1][col1-1] - this.sums[row2][col1-1] - this.sums[row1-1][col2]
	}
}