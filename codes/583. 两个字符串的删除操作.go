// 动态规划
func min(a int, b int, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	ans := make([][]int, m+1)
	ans[0] = make([]int, n+1)
	for j := 0; j <= n; j++ {
		ans[0][j] = j
	}
	for i := 1; i <= m; i++ {
		ans[i] = make([]int, n+1)
		ans[i][0] = i
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				ans[i][j] = ans[i-1][j-1]
			} else {
				ans[i][j] = min(ans[i-1][j]+1, ans[i][j-1]+1, ans[i-1][j-1]+2)
			}
		}
	}
	return ans[m][n]
}