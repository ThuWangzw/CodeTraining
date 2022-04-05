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

func minMoves(nums []int, limit int) int {
	moves := make([]int, limit*2+4)
	n := len(nums)
	for i := 0; i < n/2; i++ {
		moves[2] += 2
		l := 1 + min(nums[i], nums[n-1-i])
		r := limit + max(nums[i], nums[n-1-i])
		moves[l] -= 1
		moves[r+1] += 1
		s := nums[i] + nums[n-1-i]
		moves[s] -= 1
		moves[s+1] += 1
	}
	ans := math.MaxInt32
	for i := 2; i <= limit*2; i++ {
		moves[i] += moves[i-1]
		ans = min(ans, moves[i])
	}
	return ans
}