func minimumMoves(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i] == 'O' {
			if i == 0 {
				ans[i] = 0
			} else {
				ans[i] = ans[i-1]
			}
		} else {
			if i <= 2 {
				ans[i] = 1
			} else {
				ans[i] = 1 + ans[i-3]
			}
		}
	}
	return ans[n-1]
}