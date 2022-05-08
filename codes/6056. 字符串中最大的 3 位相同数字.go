func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func largestGoodInteger(num string) string {
	n := len(num)
	if n <= 2 {
		return ""
	}
	maxint := -1

	for i := 0; i < n-2; i++ {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			maxint = max(maxint, int(num[i]-'0'))
		}
	}
	if maxint == -1 {
		return ""
	} else {
		bit := strconv.Itoa(maxint)
		return bit + bit + bit
	}
}