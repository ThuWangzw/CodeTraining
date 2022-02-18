func qsort(nums []int, begin, end int) {
	if begin >= end {
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
	qsort(nums, right+1, end)
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	qsort(nums, 0, n-1)
	ans := make([][]int, 0)
	for i, a := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		sum := -a
		for j < k {
			if nums[j]+nums[k] == sum {
				ans = append(ans, []int{a, nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if nums[j]+nums[k] < sum {
				j++
			} else {
				k--
			}

		}
	}
	return ans
}