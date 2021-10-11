// 哈希表
func twoSum(nums []int, target int) []int {
	numCnt := make(map[int][]int)
	for i, num := range nums {
		_, ok := numCnt[num]
		if !ok {
			numCnt[num] = make([]int, 0)
		}
		numCnt[num] = append(numCnt[num], i)
	}
	for num, cnt := range numCnt {
		otherCnt, ok := numCnt[target-num]
		if len(cnt) > 0 && ok && (num != target-num || len(cnt) > 1) {
			return []int{cnt[0], otherCnt[len(otherCnt)-1]}
		}
	}
	return []int{0, 0}
}