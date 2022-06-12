func matchReplacement(s string, sub string, mappings [][]byte) bool {
	dict := make(map[byte]map[byte]bool)
	for _, mapping := range mappings {
		if _, ok := dict[mapping[0]]; !ok {
			dict[mapping[0]] = make(map[byte]bool)
		}
		dict[mapping[0]][mapping[1]] = true
	}
	for i := 0; i <= len(s)-len(sub); i++ {
		flag := true
		for j := 0; j < len(sub); j++ {
			if s[i+j] != sub[j] && !dict[sub[j]][s[i+j]] {
				flag = false
			}
		}
		if flag {
			return true
		}
	}
	return false
}