func subarraySum(nums []int, k int) int {
	ans := 0
	n := len(nums)
	presum := 0
	sumMap := make(map[int]int)
	sumMap[0] = 1
	for i := 0; i < n; i++ {
		presum += nums[i]
		sumMap[presum]++
		if k == 0 {
			ans += sumMap[presum-k] - 1
		} else {
			ans += sumMap[presum-k]
		}
	}
	return ans
}