func findKthLargest(nums []int, k int) int {
	left := 0
	right := len(nums) - 1
	pivot := nums[0]
	for left <= right {
		if nums[left] >= pivot {
			left++
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}
	nums[0], nums[right] = nums[right], nums[0]
	if right+1 == k {
		return nums[right]
	} else if right+1 < k {
		return findKthLargest(nums[right+1:], k-(right+1))
	} else {
		return findKthLargest(nums[:right], k)
	}
}