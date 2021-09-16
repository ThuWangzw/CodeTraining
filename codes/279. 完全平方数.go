// 动态规划
func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func numSquares(n int) int {
	counts := make([]int, n+1)
	for i := 1; i <= n; i++ {
		nsqrt := int(math.Sqrt(float64(i)))
		counts[i] = i
		for j := nsqrt; j >= 1; j-- {
			counts[i] = min(counts[i-j*j]+1, counts[i])
		}
	}
	return counts[n]
}