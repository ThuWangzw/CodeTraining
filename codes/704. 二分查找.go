// 关键在于梳理一下退出条件和边界值，不要找不到或者进入死循环
func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid_idx := (start + end) / 2
		if nums[mid_idx] == target {
			return mid_idx
		} else if nums[mid_idx] < target {
			start = mid_idx + 1
		} else {
			end = mid_idx - 1
		}
	}
	return -1
}