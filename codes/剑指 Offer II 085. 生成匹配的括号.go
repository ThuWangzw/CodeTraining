var ans []string

func generateParenthesisIter(s string, left int, right int) {
	if left == 0 && right == 0 {
		ans = append(ans, s)
	}
	if left > 0 {
		generateParenthesisIter(s+"(", left-1, right)
	}
	if right > 0 && right > left {
		generateParenthesisIter(s+")", left, right-1)
	}
}

func generateParenthesis(n int) []string {
	ans = make([]string, 0)
	generateParenthesisIter("", n, n)
	return ans
}