func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func rearrangeCharacters(s string, target string) int {
	targetCnt := make(map[rune]int)
	sCnt := make(map[rune]int)
	ans := math.MaxInt32
	for _, c := range target {
		targetCnt[c]++
	}
	for _, c := range s {
		sCnt[c]++
	}
	for ch, cnt := range targetCnt {
		ans = min(ans, sCnt[ch]/cnt)
	}
	return ans
}