func isOneBitCharacter(bits []int) bool {
	p := 0
	n := len(bits)
	if n == 1 {
		return true
	}
	for p < n {
		if bits[p] == 1 {
			p += 2
		} else {
			p += 1
		}
		if p == n-1 {
			return true
		}
	}
	return false
}