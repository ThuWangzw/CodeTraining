func lastRemaining(n int, m int) int {
	last := 0
	for i := 2; i <= n; i++ {
		last = ((m-1)%i + last + 1) % i
	}
	return last
}