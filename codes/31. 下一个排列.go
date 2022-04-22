func nextPermutation(nums []int) {
	n := len(nums)
	pivot := n - 1
	for pivot > 0 && nums[pivot] <= nums[pivot-1] {
		pivot--
	}
	if pivot == 0 {
		// the last permutation
		for i := 0; i < n/2; i++ {
			nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
		}
		return
	}
	pivot--
	nextPivotNum := math.MaxInt32
	nextPivot := -1
	for i := pivot + 1; i < n; i++ {
		if nums[i] > nums[pivot] && nums[i] <= nextPivotNum {
			nextPivotNum = nums[i]
			nextPivot = i
		}
	}
	nums[pivot], nums[nextPivot] = nums[nextPivot], nums[pivot]
	remainLen := n - pivot - 1
	remainNums := nums[pivot+1:]
	for i := 0; i < remainLen/2; i++ {
		remainNums[i], remainNums[remainLen-1-i] = remainNums[remainLen-1-i], remainNums[i]
	}
}