// 区间动态规划
func countSubstrings(s string) int {
	n := len(s)
	ans := make([][]bool, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]bool, n)
	}
	cnt := 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				ans[i][j] = true
			} else if i+1 == j {
				ans[i][j] = (s[i] == s[j])
			} else {
				ans[i][j] = (s[i] == s[j]) && ans[i+1][j-1]
			}
			if ans[i][j] {
				cnt++
			}
		}
	}
	return cnt
}