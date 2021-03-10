func climbStairs(n int) int {
	result := make([]int, n+1, n+1)
	result[0] = 1
	for i := 1; i <= n; i++ {
		if i == 1 {
			result[i] = result[i-1]
		} else {
			result[i] = result[i-1] + result[i-2]
		}
	}
	return result[n]
}