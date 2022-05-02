func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	lasta, lastb := 1, 1
	for i := 2; i <= n; i++ {
		lasta, lastb = lastb, (lasta+lastb)%(1e9+7)
	}
	return lastb
}