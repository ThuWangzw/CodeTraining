// 前缀乘积+后缀乘积+一些优化空间的小技巧
func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	last := 1
	for i := n - 1; i >= 0; i-- {
		res[i] = res[i] * last
		last *= nums[i]
	}
	return res
}