// 先找到最小值、最大值，再用二分查找进行搜索

func search(nums []int, target int) bool {
	start := 0
	end := len(nums) - 1
	mid := 0
	min_idx := -1
	for start <= end {
		if nums[start] < nums[end] {
			min_idx = start
			break
		}
		if end-start <= 1 {
			min_idx = end
			break
		}
		mid = (start + end) / 2
		if nums[mid] < nums[start] {
			end = mid
		} else if nums[mid] > nums[start] {
			start = mid
		} else {
			start++
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
			return true
		}
	}
	return false
}