func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	ans := 0
	area := make([]int, len(heights))
	for i, height := range heights {
		for len(stack) > 0 {
			n := len(stack)
			if heights[stack[n-1]] <= height {
				break
			}
			area[stack[n-1]] = heights[stack[n-1]] * (i - stack[n-1])
			stack = stack[:n-1]
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		n := len(stack)
		area[stack[n-1]] = heights[stack[n-1]] * (len(heights) - stack[n-1])
		stack = stack[:n-1]
	}
	for i := len(heights) - 1; i >= 0; i-- {
		height := heights[i]
		for len(stack) > 0 {
			n := len(stack)
			if heights[stack[n-1]] <= height {
				break
			}
			area[stack[n-1]] += heights[stack[n-1]] * (stack[n-1] - i - 1)
			stack = stack[:n-1]
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		n := len(stack)
		area[stack[n-1]] += heights[stack[n-1]] * (stack[n-1])
		stack = stack[:n-1]
	}
	for _, a := range area {
		ans = max(ans, a)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}