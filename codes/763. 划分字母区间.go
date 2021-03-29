// 贪心算法
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func partitionLabels(S string) []int {
	charRecord := make(map[byte][]int)
	for i := 'a'; i <= 'z'; i++ {
		charRecord[byte(i)] = []int{-1, -1}
	}
	n := len(S)
	for i := 0; i < n; i++ {
		if charRecord[S[i]][0] == -1 {
			charRecord[S[i]][0] = i
		}
		charRecord[S[i]][1] = i
	}
	partitions := make([]int, 0)
	begin := 0
	last := 0
	for p := 0; p < n; p++ {
		last = max(charRecord[S[p]][1], last)
		if p == last {
			partitions = append(partitions, last-begin+1)
			begin = p + 1
		}
	}
	return partitions
}