// 动态规划+简单的多个状态+时间上的优化
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	up := make([]int, n)
	down := make([]int, n)
	up[0] = 1
	down[0] = 1

	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			up[i] = up[i-1]
		} else {
			up[i] = max(up[i-1], down[i-1]+1)
		}
		if nums[i] >= nums[i-1] {
			down[i] = down[i-1]
		} else {
			down[i] = max(down[i-1], up[i-1]+1)
		}
	}
	return max(up[n-1], down[n-1])
}