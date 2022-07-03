func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	for start <= end {
		for (start <= end) && (s[start] < 'a' || s[start] > 'z') && (s[start] < 'A' || s[start] > 'Z') && (s[start] < '0' || s[start] > '9') {
			start++
		}
		for (start <= end) && (s[end] < 'a' || s[end] > 'z') && (s[end] < 'A' || s[end] > 'Z') && (s[end] < '0' || s[end] > '9') {
			end--
		}
		if start > end {
			break
		}
		sa := s[start]
		if sa >= 'A' && sa <= 'Z' {
			sa = (sa - 'A') + 'a'
		}
		sb := s[end]
		if sb >= 'A' && sb <= 'Z' {
			sb = (sb - 'A') + 'a'
		}
		if sa == sb {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}