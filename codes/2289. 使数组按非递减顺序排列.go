func totalSteps(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	steps := make([]int, n)
	stack := []int{0}
	ans := 0
	for i := 1; i < n; i++ {
		for {
			if len(stack) == 0 {
				break
			}
			top := stack[len(stack)-1]

			if nums[i] >= nums[top] {
				stack = stack[:len(stack)-1]
				steps[i] = max(steps[i], steps[top]+1)
			} else {
				break
			}
		}
		if len(stack) == 0 {
			steps[i] = 0
		} else {
			steps[i] = max(steps[i], 1)
		}
		ans = max(ans, steps[i])
		stack = append(stack, i)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}