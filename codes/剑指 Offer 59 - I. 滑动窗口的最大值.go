func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	stack := make([]int, 0)
	for i := 0; i < k; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	ans := make([]int, 0)
	ans = append(ans, nums[stack[0]])
	for j := k; j < len(nums); j++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[j] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, j)
		if stack[0] == j-k {
			stack = stack[1:]
		}
		ans = append(ans, nums[stack[0]])
	}
	return ans
}