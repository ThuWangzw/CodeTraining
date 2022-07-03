func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			ans := 0
			switch token {
			case "+":
				ans = a + b
			case "-":
				ans = a - b
			case "*":
				ans = a * b
			case "/":
				ans = a / b
			}
			stack = append(stack, ans)
		} else {
			stack = append(stack, num)
		}
	}
	return stack[0]
}