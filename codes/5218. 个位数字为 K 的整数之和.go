func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	dp := make([]int, num+1)
	for i := 0; i <= num; i++ {
		if i%10 == k {
			dp[i] = 1
		} else {
			dp[i] = math.MaxInt32
			for j := 0; j*10+k <= i; j++ {
				dp[i] = min(dp[i], dp[i-j*10-k]+1)
			}
		}
	}
	if dp[num] == math.MaxInt32 {
		return -1
	}
	return dp[num]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}