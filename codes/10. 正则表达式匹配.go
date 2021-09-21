// 动态规划，复杂的边界条件。。。
func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	ans := make([][]bool, m+1)
	ans[0] = make([]bool, n+1)
	ans[0][0] = true
	for i := 0; i <= m; i++ {
		ans[i] = make([]bool, n+1)
		ans[0][0] = true
		for j := 1; j <= n; j++ {
			if i > 0 && s[i-1] == p[j-1] {
				ans[i][j] = ans[i-1][j-1]
			} else if i > 0 && p[j-1] != '.' && p[j-1] != '*' {
				ans[i][j] = false
			} else if i > 0 && p[j-1] == '.' {
				ans[i][j] = ans[i-1][j-1]
			} else if p[j-1] == '*' && p[j-2] == '.' {
				for k := i; k >= 0; k-- {
					ans[i][j] = ans[i][j] || ans[k][j-2]
				}
			} else if p[j-1] == '*' {
				for k := i; k >= 0 && (k == i || s[k] == p[j-2]); k-- {
					ans[i][j] = ans[i][j] || ans[k][j-2]
				}
			}
		}
	}
	return ans[m][n]
}