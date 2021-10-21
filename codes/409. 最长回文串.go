func longestPalindrome(s string) int {
	cnts := make(map[byte]int)
	for _, ch := range s {
		cnts[byte(ch)]++
	}
	cnt := 0
	hasEven := false
	for _, chcnt := range cnts {
		if chcnt%2 == 0 {
			cnt += chcnt
		} else if !hasEven {
			hasEven = true
			cnt += chcnt
		} else {
			cnt += chcnt - 1
		}
	}
	return cnt
}