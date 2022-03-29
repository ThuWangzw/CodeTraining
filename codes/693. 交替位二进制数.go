func hasAlternatingBits(n int) bool {
	lastBit := n % 2
	for n > 0 {
		lastBit = n % 2
		n = n >> 1
		if n%2 == lastBit {
			return false
		}
	}
	return true
}