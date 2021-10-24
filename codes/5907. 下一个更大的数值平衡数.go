func isBalance(n int) bool {
	cnts := make(map[int]int)
	if n == 0 {
		return false
	}
	for n > 0 {
		cnts[n%10]++
		n /= 10
	}
	for k, cnt := range cnts {
		if k != cnt {
			return false
		}
	}
	return true
}

func nextBeautifulNumber(n int) int {
	n++
	for !isBalance(n) {
		n++
	}
	return n
}