func longestSubsequence(s string, k int) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	if s[n-1] == '0' {
		return longestSubsequence(s[:n-1], k>>1) + 1
	} else if k%2 == 1 {
		return longestSubsequence(s[:n-1], k>>1) + 1
	} else {
		j := n - 1
		for ; j >= 0; j-- {
			if s[j] == '0' {
				break
			}
		}
		if k == 0 {
			return longestSubsequence(s[:j+1], k>>1)
		}
		return max(longestSubsequence(s[:n-1], (k-1)>>1)+1, longestSubsequence(s[:j+1], k>>1))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}