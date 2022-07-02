func findTargetSumWays(nums []int, target int) int {
	m := len(nums)
	n := 2001
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][1000] = 1
	for i := 1; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if j-nums[i-1] >= 0 {
				dp[i][j] += dp[i-1][j-nums[i-1]]
			}
			if j+nums[i-1] < n {
				dp[i][j] += dp[i-1][j+nums[i-1]]
			}
		}
	}
	return dp[m][target+1000]
}

// dp[i][j] = dp[i-1][j-nums[i]]+dp[i-1][j+nums[i]]