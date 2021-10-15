// 哈希表
func findShortestSubArray(nums []int) int {
	cnts := make(map[int][]int)
	for i, num := range nums {
		if _, ok := cnts[num]; !ok {
			cnts[num] = make([]int, 3)
			cnts[num][1] = i
		}
		cnts[num][0]++
		cnts[num][2] = i
	}
	maxcnt := 0
	minlen := math.MaxInt32
	for _, cnt := range cnts {
		if maxcnt < cnt[0] {
			maxcnt = cnt[0]
			minlen = cnt[2] - cnt[1] + 1
		} else if maxcnt == cnt[0] {
			minlen = min(minlen, cnt[2]-cnt[1]+1)
		}
	}
	return minlen
}