func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minMoves(nums []int) int {
	minv := math.MaxInt32
	for _, num := range nums {
		minv = min(minv, num)
	}
	sum := 0
	for _, num := range nums {
		sum += num - minv
	}
	return sum
}