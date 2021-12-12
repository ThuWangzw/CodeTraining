// golang本身自带的string的功能少，效率低，可以使用strings库
func toLowerCase(s string) string {
	ans := new(strings.Builder)
	ans.Grow(len(s))
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			ans.WriteRune(ch - 'A' + 'a')
		} else {
			ans.WriteRune(ch)
		}
	}
	return ans.String()
}