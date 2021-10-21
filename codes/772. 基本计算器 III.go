// 后缀表达式转中缀
func computeOne(a int, b int, ch byte) int {
	switch ch {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	return -1
}

func computeOneFromStack(numstack []int, oprstack []byte) ([]int, []byte) {
	opr := oprstack[len(oprstack)-1]
	oprstack = oprstack[:len(oprstack)-1]

	one := numstack[len(numstack)-1]
	numstack = numstack[:len(numstack)-1]

	sec := numstack[len(numstack)-1]
	numstack = numstack[:len(numstack)-1]

	numstack = append(numstack, computeOne(sec, one, opr))
	return numstack, oprstack
}

func calculate(s string) int {
	numstack := make([]int, 0)
	oprstack := make([]byte, 0)
	cur := 0
	for i, ch := range s {
		if ch >= '0' && ch <= '9' {
			cur = cur*10 + int(byte(ch)-'0')
			continue
		} else if i > 0 && s[i-1] >= '0' && s[i-1] <= '9' {
			numstack = append(numstack, cur)
			cur = 0
		}

		if ch == '(' {
			oprstack = append(oprstack, '(')
		} else if ch == '+' || ch == '-' {
			for len(oprstack) > 0 && oprstack[len(oprstack)-1] != '(' {
				numstack, oprstack = computeOneFromStack(numstack, oprstack)
			}
			oprstack = append(oprstack, byte(ch))
		} else if ch == '*' || ch == '/' {
			for len(oprstack) > 0 && (oprstack[len(oprstack)-1] == '*' || oprstack[len(oprstack)-1] == '/') {
				numstack, oprstack = computeOneFromStack(numstack, oprstack)
			}
			oprstack = append(oprstack, byte(ch))
		} else if ch == ')' {
			for oprstack[len(oprstack)-1] != '(' {
				numstack, oprstack = computeOneFromStack(numstack, oprstack)
			}
			oprstack = oprstack[:len(oprstack)-1]
		}
	}
	if s[len(s)-1] != ')' {
		numstack = append(numstack, cur)
	}
	for len(oprstack) > 0 {
		numstack, oprstack = computeOneFromStack(numstack, oprstack)
	}
	return numstack[0]
}