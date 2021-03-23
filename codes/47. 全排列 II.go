func permuteUniqueIter(nums map[int]int) [][]int {
	results := make([][]int, 0)
	for num := range nums {
		if nums[num] > 0{
			nums[num]--
			oneResult := permuteUniqueIter(nums)
			ori := make([]int, 1)
			ori[0] = num
			for _, result := range oneResult {
				results = append(results, append(ori, result...))
			}
			nums[num]++
		}
	}
	if len(results) == 0 {
		results = append(results, make([]int, 0))
	}
	return results
}

func permuteUnique(nums []int) [][]int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	return permuteUniqueIter(numMap)
}