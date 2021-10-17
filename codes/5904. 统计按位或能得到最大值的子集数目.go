// 递归回溯
func countIter(nums []int, base int, target int, count *int) {
	if len(nums) == 0 {
		if base == target {
			*count++
		}
	} else {
		countIter(nums[1:], base, target, count)
		countIter(nums[1:], base|nums[0], target, count)
	}
}

func countMaxOrSubsets(nums []int) int {
	count := 0
	target := 0
	for _, num := range nums {
		target |= num
	}
	countIter(nums, 0, target, &count)
	return count
}