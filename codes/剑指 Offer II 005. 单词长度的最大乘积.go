func maxProduct(words []string) int {
	n := len(words)
	wordMaps := make([]int, 0, n)
	for _, word := range words {
		wordMap := 0
		for _, c := range word {
			wordMap = wordMap | (1 << int(c-'a'))
		}
		wordMaps = append(wordMaps, wordMap)
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if wordMaps[i]&wordMaps[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}