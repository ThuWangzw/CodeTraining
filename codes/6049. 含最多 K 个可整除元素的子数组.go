func isEqual(nums []int, a int, b int, s int) bool {
	for i := 0; i <= s; i++ {
		if nums[a+i] != nums[b+i] {
			return false
		}
	}
	return true
}
func countDistinct(nums []int, k int, p int) int {
	n := len(nums)
	count := 0
	for s := 0; s < n; s++ {
		pCount := 0
		for i := 0; i <= s; i++ {
			if nums[i]%p == 0 {
				pCount++
			}
		}
		if pCount <= k {
			count++
		}
		for i := 1; i < n-s; i++ {
			if nums[i-1]%p == 0 {
				pCount--
			}
			if nums[i+s]%p == 0 {
				pCount++
			}
			if pCount <= k {
				flag := false
				for j := 0; j < i; j++ {
					if isEqual(nums, j, i, s) {
						flag = true
						break
					}
				}
				if !flag {
					count++
				}
			}
		}
	}
	return count
}