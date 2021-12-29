func countQuadruplets(nums []int) int {
	cnt := 0
	n := len(nums)
	for ai := 0; ai < n; ai++ {
		for bi := ai + 1; bi < n; bi++ {
			for ci := bi + 1; ci < n; ci++ {
				for di := ci + 1; di < n; di++ {
					if nums[ai]+nums[bi]+nums[ci] == nums[di] {
						cnt++
					}
				}
			}
		}
	}
	return cnt
}