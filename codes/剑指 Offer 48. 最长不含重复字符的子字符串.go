func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	lastPos := make([]int, n)
	byteMap := make(map[byte]int)
	for i := 0; i < n; i++ {
		if pos, ok := byteMap[s[i]]; ok {
			lastPos[i] = pos
		} else {
			lastPos[i] = -1
		}
		byteMap[s[i]] = i
	}
	dp := make([]int, n)
	dp[0] = 1
	ans := 1
	for i := 1; i < n; i++ {
		dp[i] = min(i-lastPos[i], dp[i-1]+1)
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}