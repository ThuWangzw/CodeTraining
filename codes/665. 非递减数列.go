// 考虑清楚再写啊
func checkPossibility(nums []int) bool {
	n := len(nums)
	first := true
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			if (!first) || ((i < n-1 && nums[i+1] < nums[i-1]) && (i > 1 && nums[i] < nums[i-2])) {
				return false
			}
			first = false
		}
	}
	return true
}