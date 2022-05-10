func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isStraight(nums []int) bool {
	minv := 14
	maxv := 0
	zerocnt := 0
	numsMap := [14]int{}
	for _, num := range nums {
		if num == 0 {
			zerocnt++
		} else {
			if numsMap[num] == 1 {
				return false
			}
			minv = min(minv, num)
			maxv = max(maxv, num)
			numsMap[num]++
		}
	}
	if zerocnt == 5 {
		return true
	}
	return maxv-minv+1-(5-zerocnt) <= zerocnt
}