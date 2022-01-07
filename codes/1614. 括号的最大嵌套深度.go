func maxDepth(s string) int {
	leftcnt := 0
	maxleftcnt := 0
	for _, ch := range s {
		if ch == '(' {
			leftcnt++
			if leftcnt > maxleftcnt {
				maxleftcnt = leftcnt
			}
		} else if ch == ')' {
			leftcnt--
		}
	}
	return maxleftcnt
}