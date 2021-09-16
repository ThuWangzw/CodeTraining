// 动态规划
func min(a int, b int, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	maxSquareMat := make([][]int, m)
	maxSquare := 0
	for i := 0; i < m; i++ {
		maxSquareMat[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				maxSquareMat[i][j] = 0
			} else if (i == 0) || (j == 0) || (matrix[i-1][j] == '0') || (matrix[i][j-1] == '0') {
				maxSquareMat[i][j] = 1
			} else {

				maxSquareMat[i][j] = 1 + min(maxSquareMat[i-1][j-1], maxSquareMat[i][j-1], maxSquareMat[i-1][j])
			}
			maxSquare = max(maxSquare, maxSquareMat[i][j])
		}
	}
	return maxSquare * maxSquare
}