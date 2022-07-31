func minimumOperations(nums []int) int {
	numMap := make(map[int]bool)
	for _, num := range nums {
		if num > 0 {
			numMap[num] = true
		}
	}
	return len(numMap)
}