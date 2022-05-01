func appealSum(s string) int64 {
	var count int64
	n := len(s)
	lastPos := make([]int, n)
	pos := make(map[byte]int)
	for i := 0; i < n; i++ {
		if last, ok := pos[s[i]]; ok {
			lastPos[i] = last
		} else {
			lastPos[i] = -1
		}
		pos[s[i]] = i
	}
	lastCount := 0
	for i := 0; i < n; i++ {
		lastCount += i - lastPos[i]
		count += int64(lastCount)
	}
	return count
}