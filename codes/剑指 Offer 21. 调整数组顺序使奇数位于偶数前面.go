func exchange(nums []int) []int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left]%2 == 1 {
			left++
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}
	return nums
}