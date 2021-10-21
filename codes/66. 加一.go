func plusOne(digits []int) []int {
	n := len(digits)
	last := 1
	for i := n - 1; i >= 0; i-- {
		digits[i], last = (digits[i]+last)%10, (digits[i]+last)/10
	}
	if last > 0 {
		digits = append([]int{last}, digits...)
	}
	return digits
}