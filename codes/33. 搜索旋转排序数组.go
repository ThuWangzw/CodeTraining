// 学会化归，用已解决问题解决现有问题
func search(nums []int, target int) int {
	// find min_idx
	start := 0
	end := len(nums) - 1
	mid := 0
	min_idx := -1
	if nums[start] <= nums[end] {
		min_idx = start
	} else {
		for start <= end {
			if end-start <= 1 {
				min_idx = end
				break
			}
			mid = (start + end) / 2
			if nums[mid] < nums[start] {
				end = mid
			} else {
				start = mid
			}
		}
	}

	//
	n := len(nums)
	start = min_idx
	end = min_idx - 1 + n
	for start <= end {
		mid := (start + end) / 2
		if nums[mid%n] > target {
			end = mid - 1
		} else if nums[mid%n] < target {
			start = mid + 1
		} else {
			return mid % n
		}
	}
	return -1
}