func minimumTotal(triangle [][]int) int {
	path_result := make([][]int, 0, len(triangle))
	first_row := []int{triangle[0][0]}
	path_result = append(path_result, first_row)
	depth := len(triangle)
	for i := 1; i < depth; i++ {
		row := make([]int, i+1, i+1)
		row[0] = path_result[i-1][0] + triangle[i][0]
		row[i] = path_result[i-1][i-1] + triangle[i][i]
		path_result = append(path_result, row)
		for j := 1; j < i; j++ {
			row[j] = int(math.Min(float64(path_result[i-1][j]), float64(path_result[i-1][j-1]))) + triangle[i][j]
		}
	}
	min_path := math.MaxInt32
	for _, v := range path_result[depth-1] {
		min_path = int(math.Min(float64(min_path), float64(v)))
	}
	return min_path
}