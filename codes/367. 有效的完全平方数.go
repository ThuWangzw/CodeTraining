func isPerfectSquare(num int) bool {
	i := 1
	for i <= num/i {
		if i == num/i {
			break
		}
		i++
	}
	if i*(num/i) == num && i == num/i {
		return true
	}
	return false
}