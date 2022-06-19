func greatestLetter(s string) string {
	cmap := make(map[rune][]bool)
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			if _, ok := cmap['A'+c-'a']; !ok {
				cmap['A'+c-'a'] = make([]bool, 2)
			}
			cmap['A'+c-'a'][1] = true
		}
		if c >= 'A' && c <= 'Z' {
			if _, ok := cmap[c]; !ok {
				cmap[c] = make([]bool, 2)
			}
			cmap[c][0] = true
		}
	}
	for c := rune('Z'); c >= rune('A'); c-- {
		if occur, ok := cmap[c]; ok && occur[0] && occur[1] {
			return string(c)
		}
	}
	return ""
}