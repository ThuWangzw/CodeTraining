func subArrayRanges(nums []int) int64 {
	var maxsum int64
	var minsum int64
	n := len(nums)
	leftmin := make([]int, n)
	leftmax := make([]int, n)
	leftminstack := make([]int, 0)
	leftmaxstack := make([]int, 0)

	rightmin := make([]int, n)
	rightmax := make([]int, n)
	rightminstack := make([]int, 0)
	rightmaxstack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(leftminstack) > 0 && nums[leftminstack[len(leftminstack)-1]] > nums[i] {
			leftminstack = leftminstack[:len(leftminstack)-1]
		}
		if len(leftminstack) == 0 {
			leftmin[i] = i + 1
		} else {
			leftmin[i] = i - leftminstack[len(leftminstack)-1]
		}
		leftminstack = append(leftminstack, i)

		for len(leftmaxstack) > 0 && nums[leftmaxstack[len(leftmaxstack)-1]] <= nums[i] {
			leftmaxstack = leftmaxstack[:len(leftmaxstack)-1]
		}
		if len(leftmaxstack) == 0 {
			leftmax[i] = i + 1
		} else {
			leftmax[i] = i - leftmaxstack[len(leftmaxstack)-1]
		}
		leftmaxstack = append(leftmaxstack, i)

		j := n - 1 - i
		for len(rightminstack) > 0 && nums[rightminstack[len(rightminstack)-1]] >= nums[j] {
			rightminstack = rightminstack[:len(rightminstack)-1]
		}
		if len(rightminstack) == 0 {
			rightmin[j] = i + 1
		} else {
			rightmin[j] = rightminstack[len(rightminstack)-1] - j
		}
		rightminstack = append(rightminstack, j)

		for len(rightmaxstack) > 0 && nums[rightmaxstack[len(rightmaxstack)-1]] < nums[j] {
			rightmaxstack = rightmaxstack[:len(rightmaxstack)-1]
		}
		if len(rightmaxstack) == 0 {
			rightmax[j] = i + 1
		} else {
			rightmax[j] = rightmaxstack[len(rightmaxstack)-1] - j
		}
		rightmaxstack = append(rightmaxstack, j)
	}
	for i := 0; i < n; i++ {
		minsum += int64(nums[i] * leftmin[i] * rightmin[i])
		maxsum += int64(nums[i] * leftmax[i] * rightmax[i])
	}
	return maxsum - minsum
}