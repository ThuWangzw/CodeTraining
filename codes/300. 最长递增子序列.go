// 方法一：暴力动规

func lengthOfLIS(nums []int) int {
	n := len(nums)
	result := make([]int, n, n)
	result[0] = 1
	max := 1
	for i := 1; i < n; i++ {
		result[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				result[i] = int(math.Max(float64(result[j]+1), float64(result[i])))
			}
		}
		max = int(math.Max(float64(max), float64(result[i])))
	}
	return max
}

// 方法二：通过二分+动规，可以实现O(nlgn)的时间复杂度
func lengthOfLIS(nums []int) int {
	n := len(nums)
	tail := make([]int, 1, n)
	tail[0] = nums[0]
	for i := 1; i < n; i++ {
		num := nums[i]
		if tail[len(tail)-1] < num {
			tail = append(tail, num)
		} else {
			begin := 0
			end := len(tail) - 1
			idx := -1
			for begin < end {
				mid := (begin + end) / 2
				if tail[mid] < num {
					begin = mid + 1
				} else if tail[mid] > num {
					end = mid
				} else {
					idx = mid
					break
				}
			}
			if idx == -1 {
				idx = end
			}
			tail[idx] = num
		}
	}
	return len(tail)
}