func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func less(a, b string) bool {
	m := len(a)
	n := len(b)
	for i := 0; i < min(m, n); i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}
	return m < n
}

func longestWord(words []string) string {
	allWordMap := make(map[string]bool)
	wordLenMap := make([][]string, 30)
	for i := 0; i < 30; i++ {
		wordLenMap[i] = make([]string, 0)
	}
	for _, word := range words {
		wordLenMap[len(word)-1] = append(wordLenMap[len(word)-1], word)
	}
	maxlen := 0
	maxword := ""

	for _, wordsBucket := range wordLenMap {
		for _, word := range wordsBucket {
			if len(word) == 1 {
				allWordMap[word] = true
			} else {
				allWordMap[word] = allWordMap[word[:len(word)-1]]
			}
			if allWordMap[word] && (maxlen < len(word) || less(word, maxword)) {
				maxlen = len(word)
				maxword = word
			}
		}
	}
	return maxword
}