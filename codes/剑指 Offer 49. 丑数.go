func nthUglyNumber(n int) int {
	uglyNumbers := []int{1, 2, 3, 4, 5, 6}
	for i := 6; i < n; i++ {
		least := math.MaxInt32
		for _, ugly := range []int{2, 3, 5} {
			j := i - 1
			for ; j >= 0; j-- {
				if ugly*uglyNumbers[j] <= uglyNumbers[i-1] {
					break
				}
			}
			least = min(least, ugly*uglyNumbers[j+1])
		}
		uglyNumbers = append(uglyNumbers, least)
	}
	return uglyNumbers[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}