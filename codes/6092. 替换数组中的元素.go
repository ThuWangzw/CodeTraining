func arrayChange(nums []int, operations [][]int) []int {
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = i
	}
	for _, operation := range operations {
		idx := numsMap[operation[0]]
		nums[idx] = operation[1]
		numsMap[operation[1]] = idx
	}
	return nums
}