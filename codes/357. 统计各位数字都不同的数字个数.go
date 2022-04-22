func countNumbersWithUniqueDigitsIter(n int, factorial10 []int, factorial9 []int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 10
	} else {
		return factorial10[n] - factorial9[n-1] + countNumbersWithUniqueDigitsIter(n-1, factorial10, factorial9)
	}
}

func countNumbersWithUniqueDigits(n int) int {
	factorial10 := make([]int, 11)
	factorial9 := make([]int, 10)
	for i := 1; i <= 10; i++ {
		if i == 1 {
			factorial10[i] = 10
		} else {
			factorial10[i] = factorial10[i-1] * (11 - i)
		}
	}
	for i := 1; i <= 9; i++ {
		if i == 1 {
			factorial9[i] = 9
		} else {
			factorial9[i] = factorial9[i-1] * (10 - i)
		}
	}
	return countNumbersWithUniqueDigitsIter(n, factorial10, factorial9)
}