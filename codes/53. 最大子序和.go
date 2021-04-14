func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	n := len(nums)
	maxsub := make([]int, n)
	maxsub[0] = nums[0]
	result := nums[0]
	for i := 1; i < n; i++ {
		maxsub[i] = max(maxsub[i-1]+nums[i], nums[i])
		result = max(result, maxsub[i])
	}
	return result
}