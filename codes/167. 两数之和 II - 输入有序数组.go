// 如果没有思路，就先排序
func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	i := 0
	j := n - 1
	for i != j {
		result := numbers[i] + numbers[j] - target
		if result == 0 {
			return []int{i + 1, j + 1}
		} else if result > 0 {
			j--
		} else {
			i++
		}
	}
	return []int{0, 0}
}