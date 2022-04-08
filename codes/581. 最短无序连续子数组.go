func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	left := 0
	for ; left < n-1; left++ {
		if nums[left] > nums[left+1] {
			break
		}
	}
	if left == n-1 {
		return 0
	}
	right := n - 1
	for ; right >= 1; right-- {
		if nums[right] < nums[right-1] {
			break
		}
	}
	minnum := math.MaxInt32
	maxnum := math.MinInt32
	for i := left; i <= right; i++ {
		minnum = min(minnum, nums[i])
		maxnum = max(maxnum, nums[i])
	}
	for left >= 0 && nums[left] > minnum {
		left--
	}
	for right < n && nums[right] < maxnum {
		right++
	}
	return right - left - 1
}