func firstBadVersion(n int) int {
	start := 1
	end := n
	for start <= end {
		mid := (start + end) / 2
		if isBadVersion(mid) {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}