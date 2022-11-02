func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var a1, a2, b1, b2 int
	for {
		if len(word1) == a1 {
			return len(word2) == b1
		}
		if len(word2) == b1 {
			return len(word1) == a1
		}
		if word1[a1][a2] != word2[b1][b2] {
			return false
		}
		a2++
		b2++
		if a2 == len(word1[a1]) {
			a1, a2 = a1+1, 0
		}
		if b2 == len(word2[b1]) {
			b1, b2 = b1+1, 0
		}
	}
}