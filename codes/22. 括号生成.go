// 递归
var res []string

func generateParenthesisIter(s string, remainLeft int, remainRight int) {
	if remainLeft == 0 && remainRight == 0 {
		res = append(res, s)
		return
	}
	if remainLeft > 0 {
		generateParenthesisIter(s+"(", remainLeft-1, remainRight)
	}
	if remainLeft < remainRight {
		generateParenthesisIter(s+")", remainLeft, remainRight-1)
	}
}

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	generateParenthesisIter("", n, n)
	return res
}