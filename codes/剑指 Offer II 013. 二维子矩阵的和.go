type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			top := 0
			left := 0
			topleft := 0
			if i > 0 {
				top = matrix[i-1][j]
			}
			if j > 0 {
				left = matrix[i][j-1]
			}
			if i > 0 && j > 0 {
				topleft = matrix[i-1][j-1]
			}
			matrix[i][j] = matrix[i][j] + top + left - topleft
		}
	}
	return NumMatrix{
		matrix: matrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	left := 0
	top := 0
	topleft := 0
	if row1 > 0 {
		top = this.matrix[row1-1][col2]
	}
	if col1 > 0 {
		left = this.matrix[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		topleft = this.matrix[row1-1][col1-1]
	}
	return this.matrix[row2][col2] - top - left + topleft
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */