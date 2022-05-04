func translateNum(num int) int {
	nums := make([]int, 0)
	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	if len(nums) == 0 {
		nums = append(nums, 0)
	}
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
	ans := make([]int, n+1)
	ans[0] = 1
	ans[1] = 1
	for i := 2; i <= n; i++ {
		ans[i] = ans[i-1]
		if nums[i-2] > 0 && nums[i-2]*10+nums[i-1] < 26 {
			ans[i] += ans[i-2]
		}
	}
	return ans[n]
}