func squareOfNum(n int) int {
	cnt := 0
	for n > 0 {
		cnt += (n % 10) * (n % 10)
		n /= 10
	}
	return cnt
}

func isHappy(n int) bool {
	isvisit := make(map[int]bool)
	for n != 1 {
		isvisit[n] = true
		n = squareOfNum(n)
		if isvisit[n] {
			return false
		}
	}
	return true
}