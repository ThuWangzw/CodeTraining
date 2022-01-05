func modifyString(s string) string {
	ans := make([]rune, len(s))
	for i, ch := range s {
		if ch != '?' {
			ans[i] = ch
		} else {
			last := rune(0)
			next := rune(0)
			if i > 0 {
				last = ans[i-1]
			}
			if i < len(s)-1 {
				next = rune(s[i+1])
			}
			for ch := rune('a'); ch <= rune('z'); ch++ {
				if ch != last && ch != next {
					ans[i] = ch
					break
				}
			}
		}
	}
	return string(ans)
}