// 单调栈再扩展一下
func nextGreaterElements(nums []int) []int {
	ms := make([]int, 0)
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		for len(ms) > 0 && ms[len(ms)-1] <= num {
			ms = ms[:len(ms)-1]
		}
		ms = append(ms, num)
	}
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		for len(ms) > 0 && ms[len(ms)-1] <= num {
			ms = ms[:len(ms)-1]
		}
		if len(ms) == 0 {
			ans[i] = -1
		} else {
			ans[i] = ms[len(ms)-1]
		}
		ms = append(ms, num)
	}
	return ans
}