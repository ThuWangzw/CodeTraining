func shortestSequence(rolls []int, k int) int {
	kMap := make(map[int]bool)
	cnt := 0
	for i, roll := range rolls {
		if kMap[roll] == false {
			cnt++
			kMap[roll] = true
		}
		if cnt == k {
			return shortestSequence(rolls[i+1:], k) + 1
		}
	}
	return 1
}