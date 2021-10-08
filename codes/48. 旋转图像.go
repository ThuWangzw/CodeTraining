// 写一下旋转后的坐标
func rotate(matrix [][]int) {
	n := len(matrix) - 1
	for i := 0; i <= n/2; i++ {
		for j := 0; j <= (n-1)/2; j++ {
			matrix[i][j], matrix[j][n-i], matrix[n-i][n-j], matrix[n-j][i] = matrix[n-j][i], matrix[i][j], matrix[j][n-i], matrix[n-i][n-j]
		}
	}
}