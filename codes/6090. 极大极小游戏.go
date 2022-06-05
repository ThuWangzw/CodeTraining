func minMaxGame(nums []int) int {
	for len(nums) > 1 {
		n := len(nums) / 2
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				nums[i] = min(nums[i*2], nums[i*2+1])
			} else {
				nums[i] = max(nums[i*2], nums[i*2+1])
			}
		}
		nums = nums[:n]
	}
	return nums[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}