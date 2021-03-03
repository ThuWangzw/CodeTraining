// 二分查找的变形，当不能通过二分来缩减搜索空间时，也可以用线性
func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	mid := 0

	for start <= end {
		if nums[start] < nums[end] {
			return nums[start]
		}
		if end-start <= 1 {
			return nums[end]
		}
		mid = (start + end) / 2
		if nums[mid] < nums[start] {
			end = mid
		} else if nums[mid] > nums[start] {
			start = mid
		} else {
			start++
		}
	}
	return -1
}