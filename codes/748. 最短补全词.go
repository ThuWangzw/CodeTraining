func CountWord(word string) map[rune]int {
	wordmap := make(map[rune]int)
	for _, ch := range word {
		if ch >= 'A' && ch <= 'Z' {
			wordmap['a'+ch-'A']++
		} else if ch >= 'a' && ch <= 'z' {
			wordmap[ch]++
		}
	}
	return wordmap
}

func shortestCompletingWord(licensePlate string, words []string) string {
	licensemap := CountWord(licensePlate)
	var ans string
	for _, word := range words {
		if ans != "" && len(ans) <= len(word) {
			continue
		}
		wordmap := CountWord(word)
		success := true
		for ch, cnt := range licensemap {
			if cnt > wordmap[ch] {
				success = false
				break
			}
		}
		if success {
			ans = word
		}
	}
	return ans
}