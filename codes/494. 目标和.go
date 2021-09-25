// 背包问题
func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		target += nums[i]
		nums[i] *= 2
	}
	if target < 0 {
		return 0
	}
	ans := make([][]int, target+1)
	for i := 0; i <= target; i++ {
		ans[i] = make([]int, n)
		if nums[0] == i {
			ans[i][0] += 1
		}
		if i == 0 {
			ans[i][0] += 1
		}
		for j := 1; j < n; j++ {
			if i >= nums[j] {
				ans[i][j] = ans[i-nums[j]][j-1] + ans[i][j-1]
			} else {
				ans[i][j] = ans[i][j-1]
			}
		}
	}
	return ans[target][n-1]
}