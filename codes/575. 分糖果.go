func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func distributeCandies(candyType []int) int {
	candyMap := make(map[int]bool)
	for _, candytype := range candyType {
		candyMap[candytype] = true
	}
	return min(len(candyType)/2, len(candyMap))
}