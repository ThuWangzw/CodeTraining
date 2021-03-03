// 感觉二分查找实现最重要的是考虑终止条件、边界条件
func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	mid := 0
	if nums[start] <= nums[end] {
		return nums[start]
	}
	for start <= end {
		if end-start <= 1 {
			return nums[end]
		}
		mid = (start + end) / 2
		if nums[mid] < nums[start] {
			end = mid
		} else {
			start = mid
		}
	}
	return -1
}