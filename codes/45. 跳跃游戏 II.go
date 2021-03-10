func jump(nums []int) int {
	n := len(nums)
	jump_cnts := make([]int, n, n)
	jump_cnts[0] = 0
	// init
	for i:=1; i<n; i++ {
		jump_cnts[i] = math.MaxInt32
	}
	// 
	for i:=0; i<n; i++ {
		end := int(math.Min(float64(nums[i]+i), float64(n-1)))
		for j := i+1; j<=end; j++ {
			jump_cnts[j] = int(math.Min(float64(jump_cnts[i]+1), float64(jump_cnts[j])))
		}
	}
	return jump_cnts[n-1]
}