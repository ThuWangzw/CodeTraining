func peopleAwareOfSecret(n int, delay int, forget int) int {
	var mod int = 1e9 + 7
	dp := make([]int, forget)
	dp[0] = 1
	for i := 1; i < n; i++ {
		for j := forget - 1; j > 0; j-- {
			dp[j] = dp[j-1]
		}
		dp[0] = 0
		for j := delay; j < forget; j++ {
			dp[0] = (dp[0] + dp[j]) % mod
		}
	}
	sum := 0
	for i := 0; i < forget; i++ {
		sum = (dp[i] + sum) % mod
	}
	return sum
}