func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	n := len(nums)
	lastsum := nums[0]
	ans := nums[0]
	for i := 1; i < n; i++ {
		lastsum = max(nums[i], lastsum+nums[i])
		ans = max(lastsum, ans)
	}
	return ans
}