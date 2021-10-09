// 栈
func isValid(s string) bool {
	stack := make([]byte, 0)
	valid := map[byte]byte{'(': ')', '{': '}', '[': ']'}

	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, byte(ch))
		} else {
			if len(stack) > 0 && valid[stack[len(stack)-1]] == byte(ch) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}