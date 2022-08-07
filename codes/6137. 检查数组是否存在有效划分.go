func validPartition(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n+1)
	dp[0] = true
	for i, _ := range nums {
		dp[i+1] = (i+1 >= 2 && nums[i] == nums[i-1] && dp[i+1-2]) || (i+1 >= 3 && nums[i] == nums[i-1] && nums[i-1] == nums[i-2] && dp[i+1-3]) || (i+1 >= 3 && nums[i] == nums[i-1]+1 && nums[i-1] == nums[i-2]+1 && dp[i+1-3])
	}
	return dp[n]
}