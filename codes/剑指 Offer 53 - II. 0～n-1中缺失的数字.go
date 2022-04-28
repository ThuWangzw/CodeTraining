func missingNumber(nums []int) int {
	begin := 0
	end := len(nums) - 1
	for begin < end {
		mid := (begin + end) / 2
		if nums[mid] != mid {
			end = mid
		} else {
			begin = mid + 1
		}
	}
	if nums[end] == end {
		return len(nums)
	}
	return end
}