func countBinarySubstrings(s string) int {
	last := 0
	cur := 0
	cnt := 0
	n := len(s)
	for i := 0; i < n; i++ {
		cur++
		if i == n-1 || s[i] != s[i+1] {
			minv := min(last, cur)
			cnt += minv
			last = cur
			cur = 0
		}
	}
	return cnt
}