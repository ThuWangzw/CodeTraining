func subsets(nums []int) [][]int {
	n := len(nums)
	if n==0 {
		return make([][]int, 0)
	} else if n==1 {
		return [][]int{{}, {nums[0]}}
	} else {
		results := subsets(nums[1:])
		m := len(results)
		ori := make([]int, 1)
		ori[0] = nums[0]
		for i:=0; i<m; i++ {
			results = append(results, append(ori, results[i]...))
		}
		return results
	}
}