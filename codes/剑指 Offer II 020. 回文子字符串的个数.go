func countIt(s *string, b, e int) int {
	if b < 0 || e >= len(*s) || (*s)[b] != (*s)[e] {
		return 0
	}
	return 1 + countIt(s, b-1, e+1)
}

func countSubstrings(s string) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		cnt += countIt(&s, i, i)
	}
	for i := 0; i < len(s)-1; i++ {
		cnt += countIt(&s, i, i+1)
	}
	return cnt
}