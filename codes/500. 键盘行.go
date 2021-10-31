func checkWord(word string, dict map[rune]int) bool {
	linecnt := -1
	for _, ch := range word {
		if linecnt == -1 {
			linecnt = dict[ch]
		}
		if linecnt != dict[ch] {
			return false
		}
	}
	return true
}

func findWords(words []string) []string {
	ch2line := make(map[rune]int)
	for _, ch := range "qwertyuiop" + "QWERTYUIOP" {
		ch2line[ch] = 0
	}
	for _, ch := range "asdfghjkl" + "ASDFGHJKL" {
		ch2line[ch] = 1
	}
	for _, ch := range "zxcvbnm" + "ZXCVBNM" {
		ch2line[ch] = 2
	}
	ans := make([]string, 0)
	for _, word := range words {
		if checkWord(word, ch2line) {
			ans = append(ans, word)
		}
	}
	return ans
}