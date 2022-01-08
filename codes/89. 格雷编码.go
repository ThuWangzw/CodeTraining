func grayCode(n int) []int {
	if n == 1 {
		return []int{0, 1}
	} else {
		ans := grayCode(n - 1)
		base := 1 << (n - 1)
		for i := base - 1; i >= 0; i-- {
			ans = append(ans, ans[i]+base)
		}
		return ans
	}
}