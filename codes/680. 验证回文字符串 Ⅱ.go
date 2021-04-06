func validPalindrome(s string) bool {
	n := len(s)
	i := 0
	j := n - 1
	first := true
	lasti := -1
	lastj := -1
	ok := true
	for i <= j {
		if s[i] != s[j] {
			if !first {
				ok = false
				break
			}
			lasti, lastj = i, j
			i++
			first = false
			continue
		}
		i++
		j--
	}
	if !ok {
		ok = true
		i, j = lasti, lastj
		j--
		for i <= j {
			if s[i] != s[j] {
				ok = false
				break
			}
			i++
			j--
		}
	}
	return ok
}