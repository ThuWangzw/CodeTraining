func luckyNumbers(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	ans := make([]int, 0)
	rowMin := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] < matrix[i][rowMin[i]] {
				rowMin[i] = j
			}
		}
	}
	for i := 0; i < m; i++ {
		flag := true
		for j := 0; j < m; j++ {
			if matrix[i][rowMin[i]] < matrix[j][rowMin[i]] {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, matrix[i][rowMin[i]])
		}
	}
	return ans
}