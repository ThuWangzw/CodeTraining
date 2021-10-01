// 首先发现是中位数，然后可以用快速选择算法求中位数
func fastSelect(nums []int, k int) int {
	n := len(nums)
	pivot := nums[0]
	start := 1
	end := n - 1
	for start <= end {
		if nums[start] > pivot {
			nums[start], nums[end] = nums[end], nums[start]
			end--
		} else {
			start++
		}
	}
	nums[0], nums[end] = nums[end], nums[0]
	if k == end+1 {
		return nums[end]
	} else if k < end+1 {
		return fastSelect(nums[:end], k)
	} else {
		return fastSelect(nums[start:], k-end-1)
	}
}

func minMoves2(nums []int) int {
	n := len(nums)
	mid := fastSelect(nums, (n+1)/2)
	cnt := 0
	for _, num := range nums {
		if num > mid {
			cnt += num - mid
		} else {
			cnt += mid - num
		}
	}
	return cnt
}