func findNthDigit(n int) int {
	k := 0
	m := 0
	lastm := 0
	for n > m {
		k += 1
		lastm = m
		m += k * 9 * int(math.Pow(10, float64(k)-1))
	}
	lastm++
	base := int(math.Pow(10, float64(k)-1))
	p := (n-lastm)/k + base
	q := k - 1 - ((n - lastm) % k)
	for q > 0 {
		p /= 10
		q--
	}
	return p % 10
}