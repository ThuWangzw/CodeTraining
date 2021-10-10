func minOperations(grid [][]int, x int) int {
	m := len(grid)
	n := len(grid[0])
	nums := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			nums[i*n+j] = grid[i][j]
		}
	}
	n = m * n
	for i := 0; i < n-1; i++ {
		if (nums[i]-nums[i+1])%x != 0 {
			return -1
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	mid := nums[n/2]
	cnt := 0
	for _, num := range nums {
		cnt += int(math.Abs(float64(num-mid))) / x
	}
	return cnt
}