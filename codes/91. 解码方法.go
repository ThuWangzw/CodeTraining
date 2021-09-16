//动态规划
func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	last := 1
	lastTwo := 1
	for i := 1; i < n; i++ {
		lastBak := last
		lastTwoBak := lastTwo
		lastTwoStr, _ := strconv.Atoi(s[i-1 : i+1])
		if (s[i] != '0') && (lastTwoStr > 9 && lastTwoStr < 27) {
			last += lastTwo
		} else if s[i] != '0' {
			last = lastBak
		} else if lastTwoStr > 9 && lastTwoStr < 27 {
			last = lastTwoBak
		} else {
			return 0
		}
		lastTwo = lastBak
	}
	return last
}