// 线性的、构造结果型的，可以用回溯
func findPalindromes(s string, isPalindrome [][]bool, begin int) [][]string {
	n := len(s)
	results := make([][]string, 0)
	if n == begin {
		results = append(results, make([]string, 0))
		return results
	}
	for i := begin; i < n; i++ {
		if isPalindrome[begin][i] {
			result := make([]string, 1)
			result[0] = s[begin : i+1]
			strs := findPalindromes(s, isPalindrome, i+1)
			for _, str := range strs {
				results = append(results, append(result, str...))
			}
		}
	}
	return results
}

func partition(s string) [][]string {
	n := len(s)
	// is palindrome
	isPalindrome := make([][]bool, 0)
	for i := 0; i < n; i++ {
		isPalindrome = append(isPalindrome, make([]bool, n, n))
		isPalindrome[i][i] = true
	}
	for i := n-1; i >= 0; i-- {
		for j := i+1; j <n; j++ {
			isPalindrome[i][j] = (s[i] == s[j]) && (j-i < 2 || isPalindrome[i+1][j-1])
		}
	}
	//
	return findPalindromes(s, isPalindrome, 0)
}