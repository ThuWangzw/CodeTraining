func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countGoodRectangles(rectangles [][]int) int {
	maxsize := 0
	count := 0
	for _, rectangle := range rectangles {
		size := min(rectangle[0], rectangle[1])
		if size > maxsize {
			maxsize, count = size, 1
		} else if size == maxsize {
			count++
		}
	}
	return count
}