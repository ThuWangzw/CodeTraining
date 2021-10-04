// 做的时候考虑太复杂了，实际上是简单的数学题
func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	msum := 0
	for _, num := range rolls {
		msum += num
	}
	sum := (m+n)*mean - msum
	if (sum < n) || (sum > 6*n) {
		return []int{}
	}
	k := sum / n
	p := sum % n
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if i < p {
			ans[i] = k + 1
		} else {
			ans[i] = k
		}
	}
	return ans
}