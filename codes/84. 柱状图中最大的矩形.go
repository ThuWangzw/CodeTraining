func largestRectangleArea(heights []int) int {
	n := len(heights)
	areas := make([]int, n)
	monoStack := make([]int, 0)
	for i := 0; i < n; i++ {
		leftbound := -1
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) > 0 {
			leftbound = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
		areas[i] = (i - leftbound) * heights[i]
	}
	monoStack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		rightbound := n
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) > 0 {
			rightbound = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
		areas[i] += (rightbound - i - 1) * heights[i]
	}
	maxarea := 0
	for _, area := range areas {
		if area > maxarea {
			maxarea = area
		}
	}
	return maxarea
}