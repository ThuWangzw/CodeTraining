func qsort(intervals [][]int, begin, end int) {
	if begin >= end {
		return
	}
	pivot := intervals[begin]
	left := begin + 1
	right := end
	for left <= right {
		if intervals[left][0] <= pivot[0] {
			left++
		} else {
			intervals[left], intervals[right] = intervals[right], intervals[left]
			right--
		}
	}
	intervals[begin], intervals[right] = intervals[right], intervals[begin]
	qsort(intervals, begin, right-1)
	qsort(intervals, right+1, end)
}

func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	qsort(intervals, 0, len(intervals)-1)
	begin, end := intervals[0][0], intervals[0][1]
	for _, interval := range intervals[1:] {
		if interval[0] > end {
			ans = append(ans, []int{begin, end})
			begin = interval[0]
			end = interval[1]
		} else if interval[1] > end {
			end = interval[1]
		}
	}
	ans = append(ans, []int{begin, end})
	return ans
}