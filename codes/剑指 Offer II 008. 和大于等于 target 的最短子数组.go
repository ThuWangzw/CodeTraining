func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	n := len(nums)
	minlen := math.MaxInt32
	for right := 0; right < n; right++ {
		sum += nums[right]
		for left <= right {
			if sum >= target {
				minlen = min(minlen, right-left+1)
				sum -= nums[left]
				left++
			} else {
				break
			}
		}
	}
	if minlen == math.MaxInt32 {
		return 0
	}
	return minlen
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}