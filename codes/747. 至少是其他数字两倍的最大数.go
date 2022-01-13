func dominantIndex(nums []int) int {
	maxnum := math.MinInt32
	maxnumidx := -1
	maxnum2 := math.MinInt32
	for i, num := range nums {
		if num > maxnum {
			maxnum, maxnum2 = num, maxnum
			maxnumidx = i
		} else if num > maxnum2 {
			maxnum2 = num
		}
	}
	if maxnum < maxnum2*2 {
		return -1
	}
	return maxnumidx
}