func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximumTop(nums []int, k int) int {
	n := len(nums)
	if n == 0 || (n == 1 && k%2 == 1) {
		return -1
	}
	end := min(n-1, k-2)
	maxnum := math.MinInt32
	for i := 0; i <= end; i++ {
		maxnum = max(maxnum, nums[i])
	}
	if k < n {
		maxnum = max(maxnum, nums[k])
	}
	return maxnum
}