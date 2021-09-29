// 首先要发现戳气球可以反过来成为加气球，然后可以用区间动态规划
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxCoins(nums []int) int {
	n := len(nums)
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	ans := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		ans[i] = make([]int, n+2)
	}
	for i := n + 1; i >= 0; i-- {
		for j := i + 1; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				ans[i][j] = max(ans[i][j], nums[i]*nums[k]*nums[j]+ans[i][k]+ans[k][j])
			}
		}
	}
	return ans[0][n+1]
}