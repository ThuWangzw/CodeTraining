func searchRange(nums []int, target int) []int {
	begin := 0
	end := len(nums) - 1
	if end == -1 {
		return []int{-1, -1}
	}
	// find start
	for begin <= end {
		mid := (begin + end) / 2
		if nums[mid] < target {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	if begin >= len(nums) || nums[begin] != target {
		return []int{-1, -1}
	}
	start := begin
	end = len(nums) - 1
	// find end
	for begin <= end {
		mid := (begin + end) / 2
		if nums[mid] <= target {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return []int{start, end}
}