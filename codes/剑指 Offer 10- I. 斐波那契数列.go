func fib(n int) int {
	if n <= 1 {
		return n
	}
	lasta, lastb := 0, 1
	for i := 2; i <= n; i++ {
		lasta, lastb = lastb, (lasta+lastb)%(1e9+7)
	}
	return lastb
}