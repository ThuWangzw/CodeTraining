func mergeSort(nums []int, merged []int) int {
	if len(nums) <= 1 {
		return 0
	}
	mid := len(nums) / 2
	leftw := mergeSort(nums[:mid], merged[:mid])
	rightw := mergeSort(nums[mid:], merged[mid:])
	ans := leftw + rightw
	leftnums, rightnums := nums[:mid], nums[mid:]
	left := 0
	right := 0
	m, n := len(leftnums), len(rightnums)
	for left < m && right < n {
		if leftnums[left] <= rightnums[right] {
			merged[left+right] = leftnums[left]
			left++
		} else {
			ans += m - left
			merged[left+right] = rightnums[right]
			right++
		}
	}
	for left < m {
		merged[left+right] = leftnums[left]
		left++
	}
	for right < n {
		ans += m - left
		merged[left+right] = rightnums[right]
		right++
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = merged[i]
	}
	return ans
}

func reversePairs(nums []int) int {
	merged := make([]int, len(nums))
	ans := mergeSort(nums, merged)
	return ans
}