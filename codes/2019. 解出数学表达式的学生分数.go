// 非常典型的区间动态规划
func scoreOfStudents(s string, answers []int) int {
	// get numbers and symbols
	n := (len(s) + 1) / 2
	symbols := make([]byte, n-1)
	numbers := make([]int, n)
	symbolIdx := 0
	numberIdx := 0
	for i := 0; i < n*2-1; i++ {
		if i%2 == 0 {
			numbers[numberIdx], _ = strconv.Atoi(string(s[i]))
			numberIdx++
		} else {
			symbols[symbolIdx] = s[i]
			symbolIdx++
		}
	}
	// count answers
	answerCnts := make([]int, 1024)
	for _, answer := range answers {
		answerCnts[answer]++
	}

	// find possible answers
	pas := make([][]map[int]bool, n)
	for i := 0; i < n; i++ {
		pas[i] = make([]map[int]bool, n)
		for j := 0; j < n; j++ {
			pas[i][j] = make(map[int]bool)
		}
		pas[i][i][numbers[i]] = true
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			for k := i; k < j; k++ {
				for pal, _ := range pas[i][k] {
					for par, _ := range pas[k+1][j] {
						if (symbols[k] == '+') && (pal+par <= 1000) {
							pas[i][j][pal+par] = true
						} else if symbols[k] == '*' && pal*par <= 1000 {
							pas[i][j][pal*par] = true
						}
					}
				}
			}
		}
	}

	// calculate right answer
	mul := 1
	rightAnswer := 0
	for i, c := range symbols {
		if c == '+' {
			rightAnswer += mul * numbers[i]
			mul = 1
		} else {
			mul *= numbers[i]
		}
	}
	rightAnswer += mul * numbers[n-1]

	//
	result := 0
	for ans, cnt := range answerCnts {
		score := 0
		if ans == rightAnswer {
			score = 5
		} else if pas[0][n-1][ans] {
			score = 2
		}
		result += score * cnt
	}

	return result
}