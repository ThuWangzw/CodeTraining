// 回溯+中缀表达式求值
func calculateOne(a int, b int, ope byte) int {
	switch ope {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	}
	return -1
}

func eval(num string, operators []byte) (int, bool) {
	n := len(num)
	cur := 0
	numStack := make([]int, 0)
	opeStack := make([]byte, 0)
	for i := 0; i < n; i++ {
		cur = cur*10 + int(num[i]-'0')
		if num[i] == '0' && (i == 0 || operators[i-1] > 0) && (i < n-1 && operators[i] == 0) {
			return 0, false
		}
		if i < n-1 && operators[i] == 0 {
			continue
		} else if i < n-1 {
			numStack = append(numStack, cur)
			if operators[i] == '+' || operators[i] == '-' {
				for len(opeStack) > 0 {
					top := numStack[len(numStack)-1]
					numStack = numStack[:len(numStack)-1]
					secondTop := numStack[len(numStack)-1]
					numStack = numStack[:len(numStack)-1]
					numStack = append(numStack, calculateOne(secondTop, top, opeStack[len(opeStack)-1]))
					opeStack = opeStack[:len(opeStack)-1]
				}
			} else if operators[i] == '*' {
				for len(opeStack) > 0 && opeStack[len(opeStack)-1] == '*' {
					top := numStack[len(numStack)-1]
					numStack = numStack[:len(numStack)-1]
					secondTop := numStack[len(numStack)-1]
					numStack = numStack[:len(numStack)-1]
					numStack = append(numStack, calculateOne(secondTop, top, opeStack[len(opeStack)-1]))
					opeStack = opeStack[:len(opeStack)-1]
				}
			}
			opeStack = append(opeStack, operators[i])
			cur = 0
		} else {
			numStack = append(numStack, cur)
		}
	}
	for len(opeStack) > 0 {
		top := numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-1]
		secondTop := numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-1]
		numStack = append(numStack, calculateOne(secondTop, top, opeStack[len(opeStack)-1]))
		opeStack = opeStack[:len(opeStack)-1]
	}
	return numStack[0], true
}

func addOperatorsIter(num string, operators *[]byte, target int, ans *[]string) {
	k := len(*operators)
	n := len(num)
	if k == n-1 {
		if res, ok := eval(num, *operators); ok && res == target {
			eval(num, *operators)
			one := make([]byte, 0)
			for i := 0; i < k; i++ {
				one = append(one, num[i])
				if (*operators)[i] > 0 {
					one = append(one, (*operators)[i])
				}
			}
			one = append(one, num[n-1])
			*ans = append(*ans, string(one))
		}
	} else {
		candidates := []byte{0, '+', '-', '*'}
		for _, operator := range candidates {
			*operators = append(*operators, operator)
			addOperatorsIter(num, operators, target, ans)
			*operators = (*operators)[:k]
		}
	}
}

func addOperators(num string, target int) []string {
	operators := make([]byte, 0)
	ans := make([]string, 0)
	addOperatorsIter(num, &operators, target, &ans)
	return ans
}