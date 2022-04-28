func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	begin := 0
	end := len(nums) - 1
	for begin < end {
		mid := (begin + end) / 2
		if nums[mid] < target {
			begin = mid + 1
		} else {
			end = mid
		}
	}
	if nums[end] != target {
		return 0
	}
	target_start := end
	begin = 0
	end = len(nums) - 1
	for begin < end {
		mid := (begin + end) / 2
		if nums[mid] < target {
			begin = mid + 1
		} else if nums[mid] == target {
			begin = mid
			if nums[end] != target {
				end--
			} else {
				begin = end
			}
		} else {
			end = mid
		}
	}
	return end - target_start + 1
}