func numSubarrayProductLessThanK(nums []int, k int) int {
	ans := 0
	left := 0
	var mul int64 = 1
	for right := 0; right < len(nums); right++ {
		mul = mul * int64(nums[right])
		for left <= right {
			if mul >= int64(k) {
				mul /= int64(nums[left])
				left++
			} else {
				ans += right - left + 1
				break
			}
		}
	}
	return ans
}