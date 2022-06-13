func countBits(n int) []int {
	ans := make([]int, n+1)
	ans[0] = 0
	for i := 1; i <= n; i++ {
		ans[i] = ans[i/2]
		if i%2 > 0 {
			ans[i]++
		}
	}
	return ans
}