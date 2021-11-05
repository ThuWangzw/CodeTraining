func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
func longestSubsequence(arr []int, difference int) int {
	n := len(arr)
	longest := make([]int, n)
	ans := 0
	id2idx := make(map[int]int)
	for i, num := range arr {
		if idx, ok := id2idx[num-difference]; ok {
			longest[i] = longest[idx] + 1
		} else {
			longest[i] = 1
		}
		id2idx[num] = i
		ans = max(ans, longest[i])
	}
	return ans
}