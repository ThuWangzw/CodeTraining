// 类似求和、排列一类的问题都可以使用回溯来解决
func combinationSum(candidates []int, target int) [][]int {
	results := make([][]int, 0)
	if len(candidates) == 0 {
		return results
	}
	n := candidates[0]
	for i := 0; target-i*n >= 0; i++ {
		ary := make([]int, i)
		for j := 0; j < i; j++ {
			ary[j] = n
		}
		if target-i*n == 0 {
			results = append(results, ary)
		} else {
			remain_result := combinationSum(candidates[1:], target-i*n)
			for _, result := range remain_result {
				results = append(results, append(ary, result...))
			}
		}
	}

	return results
}