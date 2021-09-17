// 背包问题（动态规划，注意空间压缩）
func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2
	res := make([]bool, sum+1)
	res[0] = true
	for j := 0; j < n; j++ {
		for i := sum; i >= 1; i-- {
			res[i] = res[i] || (i >= nums[j] && res[i-nums[j]])
		}
	}
	return res[sum]
}