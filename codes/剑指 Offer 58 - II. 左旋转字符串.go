func reverseLeftWords(s string, n int) string {
	if len(s) == 0 {
		return s
	}
	n = n % len(s)
	return s[n:] + s[:n]
}