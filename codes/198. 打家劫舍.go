// 动态规划
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	maxRob := make([]int, n)
	maxRob[0] = nums[0]
	maxRob[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		maxRob[i] = max(maxRob[i-1], maxRob[i-2]+nums[i])
	}
	return maxRob[n-1]
}