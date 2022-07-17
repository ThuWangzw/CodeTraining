func numberOfPairs(nums []int) []int {
	paircnt := 0
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	for _, cnt := range numMap {
		paircnt += cnt / 2
	}
	return []int{paircnt, len(nums) - paircnt*2}
}