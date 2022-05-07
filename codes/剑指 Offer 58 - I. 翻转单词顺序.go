func reverseWords(s string) string {
	ans := strings.Builder{}
	right := len(s) - 1
	for right >= 0 {
		if s[right] == ' ' {
			right--
			continue
		}
		left := right
		for left >= 0 && s[left] != ' ' {
			left--
		}
		if ans.Len() > 0 {
			ans.WriteByte(' ')
		}
		ans.WriteString(s[left+1 : right+1])
		right = left
	}
	return ans.String()
}