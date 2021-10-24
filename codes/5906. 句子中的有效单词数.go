func countValidWords(sentence string) int {
	ans := 0
	tokenlen := 0
	hasend := false
	ok := true
	hasmid := false
	hasmidright := false
	for _, ch := range sentence {
		if ch == ' ' {
			if ok && tokenlen > 0 && (!hasmid || hasmidright) {
				ans++
			}
			tokenlen = 0
			hasend = false
			ok = true
			hasmid = false
			hasmidright = false
		} else if !ok {
			continue
		} else if hasend {
			ok = false
		} else if ch == '!' || ch == '.' || ch == ',' {
			hasend = true
			tokenlen++
		} else if hasmid && ch == '-' {
			ok = false
		} else if ch == '-' && tokenlen == 0 {
			ok = false
		} else if ch >= 'a' && ch <= 'z' {
			tokenlen++
			if hasmid {
				hasmidright = true
			}
		} else if ch >= '0' && ch <= '9' {
			ok = false
		} else if ch == '-' {
			hasmid = true
		}
	}
	if ok && tokenlen > 0 {
		ans++
	}
	return ans
}