// 用位运算记录是否存在
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProduct(words []string) int {
	n := len(words)
	wordsAlpha := make([]int, n)
	for i, word := range words {
		alpha := 0
		for c := range word {
			alpha |= 1 << int(c-'a')
		}
		wordsAlpha[i] = alpha
	}

	ans := 0
	for i, a := range wordsAlpha {
		for j := i + 1; j < n; j++ {
			b := wordsAlpha[j]
			if a&b == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}