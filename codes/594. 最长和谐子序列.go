// 哈希表
func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func findLHS(nums []int) int {
	cnts := make(map[int]int)
	maxlen := 0
	for _, num := range nums {
		cnts[num]++
		if cnts[num+1] > 0 || cnts[num-1] > 0 {
			maxlen = max(maxlen, cnts[num]+cnts[num+1])
			maxlen = max(maxlen, cnts[num]+cnts[num-1])
		}
	}
	return maxlen
}