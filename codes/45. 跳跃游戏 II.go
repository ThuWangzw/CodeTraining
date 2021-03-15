// 贪心算法
func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	step := 0
	maxPos := 0
	end := 0
	for i := 0; i <= end; i++ {
		maxPos = int(math.Max(float64(maxPos), float64(i+nums[i])))
		if maxPos >= n-1 {
			return step + 1
		}
		if end == i {
			step++
			end = maxPos
		}
	}
	return -1
}