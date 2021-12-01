func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxPower(s string) int {
	ans := 1
	var last rune
	cnt := 0
	for _, ch := range s {
		if ch == last {
			cnt++
		} else {
			cnt = 1
			last = ch
		}
		ans = max(cnt, ans)
	}
	return ans
}