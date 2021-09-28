// 分治法，分成奇偶数
func beautifulArray(n int) []int {
	ans := make([][]int, n)
	ans[0] = []int{1}
	for i := 2; i <= n; i++ {
		ans[i-1] = make([]int, 0)
		atop := 0
		btop := 0
		if i%2 == 0 {
			atop = i / 2
			btop = i / 2
		} else {
			atop = (i - 1) / 2
			btop = (i + 1) / 2
		}
		ans[i-1] = append(ans[atop-1], ans[btop-1]...)
		for j := 0; j < atop; j++ {
			ans[i-1][j] *= 2
		}
		for j := atop; j < i; j++ {
			ans[i-1][j] = ans[i-1][j]*2 - 1
		}
	}
	return ans[n-1]
}