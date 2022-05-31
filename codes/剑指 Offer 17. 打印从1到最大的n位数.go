func printNumbers(n int) []int {
	maxnbit := int(math.Pow(10, float64(n))) - 1
	ans := make([]int, 0, maxnbit)
	for i := 1; i <= maxnbit; i++ {
		ans = append(ans, i)
	}
	return ans
}