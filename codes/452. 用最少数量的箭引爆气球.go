type Intervals [][]int

func (intervals Intervals) Len() int {
	return len(intervals)
}
func (intervals Intervals) Swap(i int, j int) {
	intervals[j], intervals[i] = intervals[i], intervals[j]
}
func (intervals Intervals) Less(i, j int) bool {
	return intervals[i][1] < intervals[j][1]
}

func findMinArrowShots(intervals [][]int) int {
	sort.Sort(Intervals(intervals))
	cnt := 0
	begin := 0
	lastMin := math.MinInt64
	n := len(intervals)
	for {
		find := false
		for i:=begin; i<n; i++ {
			if intervals[i][0] > lastMin {
				find = true
				begin = i
				lastMin = intervals[i][1]
				cnt++
			}
		}
		if !find {
			break
		}
	}
	return cnt
}