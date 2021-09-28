// 区间动态规划
func diffWaysToCompute(expression string) []int {
	// parse data
	numbers := make([]int, 0)
	operations := make([]byte, 0)
	num := 0
	for _, c := range expression {
		if c == '+' || c == '-' || c == '*' {
			operations = append(operations, byte(c))
			numbers = append(numbers, num)
			num = 0
		} else {
			tmp, _ := strconv.Atoi(string(c))
			num = num*10 + tmp
		}
	}
	numbers = append(numbers, num)

	n := len(numbers)
	pa := make([][][]int, n)
	for i := 0; i < n; i++ {
		pa[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			pa[i][j] = make([]int, 0)
		}
		pa[i][i] = append(pa[i][i], numbers[i])
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			for k := i; k < j; k++ {
				for _, lnum := range pa[i][k] {
					for _, rnum := range pa[k+1][j] {
						switch operations[k] {
						case '+':
							pa[i][j] = append(pa[i][j], lnum+rnum)
						case '-':
							pa[i][j] = append(pa[i][j], lnum-rnum)
						case '*':
							pa[i][j] = append(pa[i][j], lnum*rnum)
						}
					}
				}
			}
		}
	}
	return pa[0][n-1]
}