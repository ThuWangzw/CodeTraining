func moveZeroes(nums []int) {
	zeroIdx := 0
	for p, num := range nums {
		if num != 0 {
			if zeroIdx != p {
				nums[zeroIdx] = num
			}
			zeroIdx++
		}
	}
	for zeroIdx < len(nums) {
		nums[zeroIdx] = 0
		zeroIdx++
	}
}