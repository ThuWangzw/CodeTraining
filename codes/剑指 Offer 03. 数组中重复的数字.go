func findRepeatNumber(nums []int) int {
	n := len(nums)
	for _, num := range nums {
		if num < 0 {
			num += n
		}
		if nums[num] < 0 {
			return num
		}
		nums[num] -= n
	}
	return -1
}