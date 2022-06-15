func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i, num := range nums {
		if i > 0 && nums[i-1] == num {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right]
			if sum < (-num) {
				left++
			} else if sum > (-num) {
				right--
			} else {
				nextAns := []int{num, nums[left], nums[right]}
				if len(ans) > 0 {
					lastAns := ans[len(ans)-1]
					if lastAns[0] != nextAns[0] || lastAns[1] != nextAns[1] || lastAns[2] != nextAns[2] {
						ans = append(ans, nextAns)
					}
				} else {
					ans = append(ans, nextAns)
				}
				left++
				right--
			}
		}
	}
	return ans
}