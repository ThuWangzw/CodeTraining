func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxNumberOfBalloons(text string) int {
	balloon := map[byte]int{'b': 1, 'a': 1, 'l': 2, 'o': 2, 'n': 1}
	counts := make(map[byte]int)
	for _, chrune := range text {
		if _, ok := balloon[byte(chrune)]; ok {
			counts[byte(chrune)]++
		}
	}
	mincnt := math.MaxInt32
	for ch, cnt := range balloon {
		mincnt = min(mincnt, counts[ch]/cnt)
	}
	return mincnt
}