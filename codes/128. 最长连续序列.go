// 使用哈希表记录一下就行
func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func longestConsecutive(nums []int) int {
	contain := make(map[int]bool)
	visited := make(map[int]bool)
	longest := make(map[int]int)
	maxlen := 0
	for _, num := range nums {
		contain[num] = true
	}
	for _, num := range nums {
		if visited[num] {
			continue
		}
		i := num
		for contain[i] {
			if !visited[i] {
				visited[i] = true
				i++
			} else {
				i += longest[i]
			}
		}
		longest[num] = i - num
		maxlen = max(maxlen, longest[num])
	}
	return maxlen
}