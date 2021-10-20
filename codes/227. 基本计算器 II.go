// 数字栈+操作符栈
func computeOne(a int, b int, opr byte) int {
	switch opr {
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

func calculate(s string) int {
	numstack := make([]int, 0)
	oprstack := make([]byte, 0)
	cur := 0
	for _, ch := range s {
		if ch == ' ' {
			continue
		} else if ch >= '0' && ch <= '9' {
			cur = cur*10 + int(ch-'0')
		} else {
			numstack = append(numstack, cur)
			cur = 0
			if ch == '*' || ch == '/' {
				for len(oprstack) > 0 && (oprstack[len(oprstack)-1] == '*' || oprstack[len(oprstack)-1] == '/') {
					opr := oprstack[len(oprstack)-1]
					oprstack = oprstack[:len(oprstack)-1]
					sec := numstack[len(numstack)-1]
					numstack = numstack[:len(numstack)-1]
					one := numstack[len(numstack)-1]
					numstack = numstack[:len(numstack)-1]
					numstack = append(numstack, computeOne(one, sec, opr))
				}
			} else {
				for len(oprstack) > 0 {
					opr := oprstack[len(oprstack)-1]
					oprstack = oprstack[:len(oprstack)-1]
					sec := numstack[len(numstack)-1]
					numstack = numstack[:len(numstack)-1]
					one := numstack[len(numstack)-1]
					numstack = numstack[:len(numstack)-1]
					numstack = append(numstack, computeOne(one, sec, opr))
				}
			}
			oprstack = append(oprstack, byte(ch))
		}
	}
	numstack = append(numstack, cur)
	for len(oprstack) > 0 {
		opr := oprstack[len(oprstack)-1]
		oprstack = oprstack[:len(oprstack)-1]
		sec := numstack[len(numstack)-1]
		numstack = numstack[:len(numstack)-1]
		one := numstack[len(numstack)-1]
		numstack = numstack[:len(numstack)-1]
		numstack = append(numstack, computeOne(one, sec, opr))
	}
	return numstack[0]
}