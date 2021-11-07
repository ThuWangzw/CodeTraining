func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
func countVowels(word string) int64 {
	var ans int64
	n := len(word)
	targets := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	for i, ch := range word {
		if _, ok := targets[ch]; ok {
			j := min(i+1, n-i)
			ans += int64(j) * int64(n-j+1)
		}
	}
	return ans
}