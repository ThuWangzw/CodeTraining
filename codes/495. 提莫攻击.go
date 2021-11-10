func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	ans := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		ans += min(timeSeries[i+1]-timeSeries[i], duration)
	}
	return ans + duration
}