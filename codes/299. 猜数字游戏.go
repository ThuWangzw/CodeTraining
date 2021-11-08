func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func getHint(secret string, guess string) string {
	n := len(secret)
	secMap := make(map[byte]int)
	guessMap := make(map[byte]int)
	bulls := 0
	cows := 0
	for i := 0; i < n; i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			secMap[secret[i]]++
			guessMap[guess[i]]++
		}
	}
	for i := 0; i < 10; i++ {
		cows += min(secMap[byte(i)+'0'], guessMap[byte(i)+'0'])
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}