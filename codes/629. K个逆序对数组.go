// 动态规划+状态压缩
func kInversePairs(n int, k int) int {
	mod := 1000000007
	ans := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		ans[i] = make([]int, k+1)
		ans[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if j >= i {
				ans[i][j] = (ans[i][j-1] + ans[i-1][j] - ans[i-1][j-i])
			} else {
				ans[i][j] = (ans[i][j-1] + ans[i-1][j])
			}
			if ans[i][j] >= mod {
				ans[i][j] -= mod
			} else if ans[i][j] < 0 {
				ans[i][j] += mod
			}
		}
	}
	return ans[n][k]
}