// 动态规划，只不过多了对最后一家的判断
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func _rob(nums []int) int {
	n := len(nums)
	last := 0
	lastlast := 0
	for i := 0; i < n; i++ {
		cur := max(last, lastlast+nums[i])
		lastlast = last
		last = cur
	}
	return last
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 2 {
		return max(nums[0], nums[1])
	} else if n == 1 {
		return nums[0]
	}
	return max(_rob(nums[1:n-2])+nums[n-1], _rob(nums[:n-1]))
}