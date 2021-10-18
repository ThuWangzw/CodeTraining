// 动态规划，记录一下每个素数乘以哪个数最小
func nthSuperUglyNumber(n int, primes []int) int {
	ans := make([]int, n)
	m := len(primes)
	ans[0] = 1
	minpoints := make([]int, m)
	for i := 0; i < m; i++ {
		minpoints[i] = 0
	}

	for i := 1; i < n; i++ {
		minv := math.MaxInt32
		minpidx := -1
		for j := 0; j < m; j++ {
			for primes[j]*ans[minpoints[j]] <= ans[i-1] {
				minpoints[j]++
			}
			if minv > primes[j]*ans[minpoints[j]] {
				minv = primes[j] * ans[minpoints[j]]
				minpidx = j
			}
		}
		ans[i] = minv
		minpoints[minpidx]++

	}
	return ans[n-1]
}