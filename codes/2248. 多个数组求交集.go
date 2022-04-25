func intersection(nums [][]int) []int {
	n := len(nums)
	numCounts := make([]int, 1001)
	for _, row := range nums {
		for _, num := range row {
			numCounts[num]++
		}
	}
	ans := make([]int, 0)
	for i := 1; i <= 1000; i++ {
		if numCounts[i] == n {
			ans = append(ans, i)
		}
	}
	return ans
}