// 动态规划
func longestPalindrome(s string) string {
	n := len(s)
	maxlen := 0
	maxi := 0
	maxj := 0
	isPalin := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPalin[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				isPalin[i][j] = true
			} else if i+1 == j {
				isPalin[i][j] = s[i] == s[j]
			} else {
				isPalin[i][j] = s[i] == s[j] && isPalin[i+1][j-1]
			}
			if isPalin[i][j] && maxlen < j-i+1 {
				maxlen = j - i + 1
				maxi = i
				maxj = j
			}
		}
	}
	return s[maxi : maxj+1]
}