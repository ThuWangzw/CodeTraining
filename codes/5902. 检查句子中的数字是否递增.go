func areNumbersAscending(s string) bool {
	last := 0
	cur := 0
	isnum := false
	for _, ch := range s {
		if ch == ' ' {
			if isnum {
				if last >= cur {
					return false
				}
				last = cur
				cur = 0
			}
		} else if ch < '0' || ch > '9' {
			isnum = false
		} else {
			isnum = true
			cur = cur*10 + int(ch-'0')
		}
	}
	if isnum {
		if last >= cur {
			return false
		}
	}
	return true
}