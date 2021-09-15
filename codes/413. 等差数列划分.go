//动态规划
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	counts := make([]int, n)
	count := 0
	for i := 2; i < n; i++ {
		delta := nums[i] - nums[i-1]
		curDelta := nums[i-1] - nums[i-2]
		if delta == curDelta {
			counts[i] = counts[i-1] + 1
		}
		count += counts[i]
	}
	return count
}