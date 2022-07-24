func repeatedCharacter(s string) byte {
	cmap := make(map[rune]bool)
	for _, c := range s {
		if cmap[c] {
			return byte(c)
		}
		cmap[c] = true
	}
	return ' '
}