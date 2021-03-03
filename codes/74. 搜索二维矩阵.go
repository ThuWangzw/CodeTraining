// 和一维的搜索没有任何的区别
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	start := 0
	end := m*n - 1
	for start <= end {
		mid := (start + end) / 2
		num := matrix[mid/n][mid%n]
		if num == target {
			return true
		} else if num < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}