// 动态规划
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	maxArea := make([]int, n)
	maxArea[0] = 0
	maxHeight := height[0]
	for i := 1; i < n; i++ {
		j := i - 1
		area := 0
		if maxHeight > height[i] {
			for j = i - 1; j >= 0 && height[j] < height[i]; j-- {
				area += height[i] - height[j]
			}
			maxArea[i] = area + maxArea[j]
		} else {
			for height[j] != maxHeight {
				area += maxHeight - height[j]
				j--
			}
		}
		maxArea[i] = area + maxArea[j]
		maxHeight = max(maxHeight, height[i])
	}
	return maxArea[n-1]
}