func mergeSort(intervals [][]int, results [][]int) {
	if len(intervals) == 0 {
		return
	}
	if len(intervals) == 1 {
		return
	}
	mid := len(intervals) / 2
	left := intervals[:mid]
	right := intervals[mid:]
	mergeSort(left, results[:mid])
	mergeSort(right, results[mid:])
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i][0] < right[j][0] {
			results[i+j] = left[i]
			i++
		} else {
			results[i+j] = right[j]
			j++
		}
	}
	for i < len(left) {
		results[i+j] = left[i]
		i++
	}
	for j < len(right) {
		results[i+j] = right[j]
		j++
	}
	for k := 0; k < len(results); k++ {
		intervals[k] = results[k]
	}
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	tmp := make([][]int, len(intervals))
	mergeSort(intervals, tmp)
	start := intervals[0][0]
	end := intervals[0][1]
	ans := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			ans = append(ans, []int{start, end})
			start = intervals[i][0]
		}
		end = max(end, intervals[i][1])
	}
	ans = append(ans, []int{start, end})
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}