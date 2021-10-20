func isIsomorphic(s string, t string) bool {
	stot := make(map[byte]byte)
	ttos := make(map[byte]byte)
	n := len(s)
	for i := 0; i < n; i++ {
		if stot[s[i]] > 0 && stot[s[i]] != t[i] {
			return false
		} else if stot[s[i]] == 0 {
			if ttos[t[i]] > 0 {
				return false
			}
			stot[s[i]] = t[i]
			ttos[t[i]] = s[i]
		}
	}
	return true
}