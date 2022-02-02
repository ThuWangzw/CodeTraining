func reversePrefix(word string, ch byte) string {
	wordAry := []byte(word)
	for idx, c := range wordAry {
		if c == ch {
			for i := 0; i <= (idx-1)/2; i++ {
				wordAry[i], wordAry[idx-i] = wordAry[idx-i], wordAry[i]
			}
			break
		}
	}
	return string(wordAry)
}