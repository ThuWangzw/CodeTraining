func constructArr(a []int) []int {
	n := len(a)
	left := make([]int, n, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			left[i] = 1
		} else {
			left[i] = left[i-1] * a[i-1]
		}
	}
	last := 1
	for i := n - 1; i >= 0; i-- {
		left[i] = left[i] * last
		last *= a[i]
	}
	return left
}