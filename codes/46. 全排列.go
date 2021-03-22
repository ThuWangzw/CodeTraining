func permute(nums []int) [][]int {
	n := len(nums)
	if n==0 {
		result := make([][]int, 1)
		result[0] = make([]int, 0)
		return result
	}
	results := make([][]int, 0)
	for i:=0; i<n; i++ {
		n_nums := make([]int, n)
		copy(n_nums, nums)
		num := n_nums[i]
		ary := make([]int, 1)
		ary[0] = num
		n_nums[i] = n_nums[0]
		n_nums[0] = num
		iresults := permute(n_nums[1:])
		m := len(iresults)
		for j:=0; j<m; j++ {
			iresults[j] = append(ary, iresults[j]...)
		}
		results = append(results, iresults...)
	}
	return results
}