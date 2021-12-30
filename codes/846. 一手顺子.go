// 排序完，再按照顺序删除即可
func qsort(nums []int, begin int, end int) {
	if end-begin <= 0 {
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

type NumCount [2]int

func isNStraightHand(hand []int, groupSize int) bool {
	n := len(hand)
	if n%groupSize > 0 {
		return false
	}
	qsort(hand, 0, n-1)
	nums := make([]NumCount, 0)
	last := -1
	for _, num := range hand {
		if num != last {
			nums = append(nums, [2]int{num, 1})
			last = num
		} else {
			nums[len(nums)-1][1]++
		}
	}
	for i := 0; i < n/groupSize; i++ {
		last := -1
		cannotcut := false
		cut := 0
		for j := 0; j < groupSize; j++ {
			if j == len(nums) {
				return false
			}
			if last != -1 && nums[j][0] != last+1 {
				return false
			}
			nums[j][1]--
			last = nums[j][0]
			if nums[j][1] != 0 {
				cannotcut = true
			}
			if nums[j][1] == 0 && cannotcut {
				return false
			}
			if nums[j][1] == 0 {
				cut = j + 1
			}
		}
		nums = nums[cut:]
	}
	return true
}