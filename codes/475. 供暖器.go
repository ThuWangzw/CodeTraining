// 排序+双指针
func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func qsort(nums []int, begin int, end int) {
	if len(nums) == 0 || begin >= end {
		return
	}
	pivot := nums[begin]
	left := begin + 1
	right := end
	for left <= right {
		if nums[left] <= pivot {
			left++
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}
	nums[begin], nums[right] = nums[right], nums[begin]
	qsort(nums, begin, right-1)
	qsort(nums, left, end)
}

func findRadius(houses []int, heaters []int) int {
	maxr := math.MinInt32
	n := len(heaters)
	leftheater := -1
	qsort(houses, 0, len(houses)-1)
	qsort(heaters, 0, len(heaters)-1)
	for _, hpos := range houses {
		for leftheater < n-1 && heaters[leftheater+1] <= hpos {
			leftheater++
		}
		if leftheater == -1 {
			maxr = max(maxr, heaters[0]-hpos)
		} else if leftheater == n-1 {
			maxr = max(maxr, hpos-heaters[leftheater])
		} else {
			maxr = max(maxr, min(heaters[leftheater+1]-hpos, hpos-heaters[leftheater]))
		}
	}
	return maxr
}