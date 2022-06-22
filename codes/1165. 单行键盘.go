func calculateTime(keyboard string, word string) int {
	kmap := make(map[rune]int)
	for i, c := range keyboard {
		kmap[c] = i
	}
	ans := 0
	for i, c := range word {
		if i == 0 {
			ans += kmap[c]
		} else {
			ans += int(math.Abs(float64(kmap[c]) - float64(kmap[rune(word[i-1])])))
		}
	}
	return ans
}