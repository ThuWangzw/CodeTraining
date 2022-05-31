func isMatch(s string, p string) bool {
	m := len(p)
	n := len(s)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if p[i-1] >= 'a' && p[i-1] <= 'z' {
				if j == 0 {
					dp[i][j] = false
				} else {
					dp[i][j] = p[i-1] == s[j-1] && dp[i-1][j-1]
				}
			} else if p[i-1] == '.' {
				if j == 0 {
					dp[i][j] = false
				} else {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				if dp[i-2][j] {
					dp[i][j] = true
				} else if j == 0 {
					dp[i][j] = false
				} else {
					dp[i][j] = (p[i-2] == '.' || (p[i-2] == s[j-1])) && dp[i][j-1]
				}
			}
		}
	}
	return dp[m][n]
}