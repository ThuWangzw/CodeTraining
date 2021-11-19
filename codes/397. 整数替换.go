func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	} else if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	} else {
		return min(integerReplacement(n-1), integerReplacement(n+1)) + 1
	}
}