func longestIdealString(s string, k int) int {
	cmap := make(map[rune]int)
	n := len(s)
	dp := make([]int, n)
	ans := 0
	for i, c := range s {
		dp[i] = 1
		for cc := c - rune(k); cc <= c+rune(k); cc++ {
			if pos, ok := cmap[cc]; ok {
				dp[i] = max(dp[i], dp[pos]+1)
			}
		}
		cmap[c] = i
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}