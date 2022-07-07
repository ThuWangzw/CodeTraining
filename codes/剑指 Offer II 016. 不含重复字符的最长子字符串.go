func lengthOfLongestSubstring(s string) int {
	cmap := make(map[rune]int)
	len := 0
	ans := 0
	for i, c := range s {
		if _, ok := cmap[c]; !ok {
			len++
		} else {
			if i-cmap[c] > len {
				len++
			} else {
				len = i - cmap[c]
			}
		}
		cmap[c] = i
		ans = max(ans, len)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}