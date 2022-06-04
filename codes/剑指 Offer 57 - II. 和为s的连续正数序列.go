func findContinuousSequence(target int) [][]int {
	ans := [][]int{}
	for k := int(math.Sqrt(float64(target*2))) + 1; k >= 2; k-- {
		sum := k * (k - 1) / 2
		if target > sum && (target-sum)%k == 0 {
			n := (target - sum) / k
			nums := []int{}
			for i := 0; i < k; i++ {
				nums = append(nums, n+i)
			}
			ans = append(ans, nums)
		}
	}
	return ans
}
