// 数学方法
func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	ans := 1
	for n > 4 {
		ans *= 3
		n -= 3
	}
	return ans * n
}