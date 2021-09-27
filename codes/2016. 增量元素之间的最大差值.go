func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maximumDifference(nums []int) int {
	n := len(nums)
	if n < 2 {
		return -1
	}
	lastMin := nums[0]
	ans := -1
	for i := 1; i < n; i++ {
		sub := nums[i] - lastMin
		if sub > 0 {
			ans = max(ans, sub)
		}
		lastMin = min(lastMin, nums[i])
	}
	return ans
}