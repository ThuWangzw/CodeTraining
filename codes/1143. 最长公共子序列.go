// 动态规划，了解最优子结构
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	lcs := make([][]int, 0, m+1)
	lcs = append(lcs, make([]int, n+1, n+1))
	for i := 1; i <= m; i++ {
		lcs = append(lcs, make([]int, n+1, n+1))
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}
	return lcs[m][n]
}