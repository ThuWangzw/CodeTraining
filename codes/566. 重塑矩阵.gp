func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])
	if m*n!=r*c {
		return mat
	}
	ans := make([][]int, r)
	for i:=0; i<r; i++ {
		ans[i] = make([]int, c)
	}
	for i:=0; i<m*n; i++ {
		ans[i/c][i%c] = mat[i/n][i%n]
	}
	return ans
}