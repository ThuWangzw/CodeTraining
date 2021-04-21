// 双指针
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func maxArea(height []int) int {
	maxarea := 0
	n := len(height)
	left := 0
	right := n - 1
	for left < right {
		maxarea = max(maxarea, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxarea
}