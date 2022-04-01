func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	n := len(nums)
	ans := make([][]int, 2)
	ans[0] = make([]int, n)
	ans[1] = make([]int, n)
	ans[0][0] = nums[0]
	ans[1][0] = nums[0]
	res := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			if ans[0][i-1] > 0 {
				ans[1][i] = ans[0][i-1] * nums[i]
				if ans[1][i-1] <= 0 {
					ans[0][i] = ans[1][i-1] * nums[i]
				} else {
					ans[0][i] = nums[i]
				}
			} else {
				ans[0][i] = ans[1][i-1] * nums[i]
				ans[1][i] = nums[i]
			}
		} else {
			if ans[0][i-1] > 0 {
				ans[0][i] = ans[0][i-1] * nums[i]
				if ans[1][i-1] <= 0 {
					ans[1][i] = ans[1][i-1] * nums[i]
				} else {
					ans[1][i] = nums[i]
				}
			} else {
				ans[0][i] = nums[i]
				ans[1][i] = ans[1][i-1] * nums[i]
			}
		}
		res = max(res, ans[0][i])
	}
	return res
}