// 一些矩阵操作
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	i := 0
	j := 0
	for i = 0; i < m; i++ {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			break
		}
	}
	i--
	if i == -1 {
		return false
	}
	for j = 0; j < n; j++ {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			break
		}
	}
	if j == n {
		return false
	}
	matrix = matrix[:i+1]
	for k := 0; k <= i; k++ {
		matrix[k] = matrix[k][j:]
	}
	return searchMatrix(matrix, target)
}