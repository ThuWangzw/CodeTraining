// 快速选择算法
func findKthLargest(nums []int, k int) int {
	return quickFind(nums, 0, len(nums)-1, k)
}

func quickFind(nums []int, begin int, end int, k int) int {
	pivot := nums[begin]
	nums[begin], nums[end] = nums[end], nums[begin]
	i := begin
	j := begin
	for j != end {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			j++
		}
	}
	if j-i+1 == k {
		return pivot
	} else if j-i+1 < k {
		return quickFind(nums, begin, i-1, k-(j-i+1))
	} else {
		return quickFind(nums, i, j, k)
	}
}