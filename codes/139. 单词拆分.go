// 典型的序列类型动态规划
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	canBreak := make([]bool, n, n)
	// construct a contain map
	containMap := make(map[string]bool)
	for _, str := range wordDict {
		containMap[str] = true
	}

	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			_, ok := containMap[s[j:i+1]]
			if ok && (j == 0 || canBreak[j-1]) {
				canBreak[i] = true
				break
			}
		}
	}
	return canBreak[n-1]
}