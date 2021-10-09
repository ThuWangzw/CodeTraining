// 检查前k个数能不能组成全排列
func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxChunksToSorted(arr []int) int {
	cnt := 0
	maxnum := 0
	for i, num := range arr {
		maxnum = max(maxnum, num)
		if maxnum == i {
			cnt++
		}
	}
	return cnt
}