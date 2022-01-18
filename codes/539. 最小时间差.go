func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func quickSort(timePoints []int, begin, end int) {
	if begin >= end {
		return
	}
	pivot := timePoints[begin]
	left := begin + 1
	right := end
	for left <= right {
		if timePoints[left] > pivot {
			timePoints[left], timePoints[right] = timePoints[right], timePoints[left]
			right--
		} else {
			left++
		}
	}
	timePoints[begin], timePoints[right] = timePoints[right], timePoints[begin]
	quickSort(timePoints, begin, right-1)
	quickSort(timePoints, right+1, end)
}

func findMinDifference(timePoints []string) int {
	n := len(timePoints)
	if n > 24*60 {
		return 0
	}
	timeInts := make([]int, n)
	for i := 0; i < n; i++ {
		h, _ := strconv.Atoi(timePoints[i][:2])
		m, _ := strconv.Atoi(timePoints[i][3:])
		timeInts[i] = h*60 + m
	}
	quickSort(timeInts, 0, n-1)
	ans := math.MaxInt32
	for i := 1; i < n; i++ {
		ans = min(ans, timeInts[i]-timeInts[i-1])
	}
	ans = min(ans, 24*60-timeInts[n-1]+timeInts[0])
	return ans
}