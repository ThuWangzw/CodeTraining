// 三个滑动窗口
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
	sum1, sum2, sum3 := 0, 0, 0
	maxsum1, maxsum2, maxsum := 0, 0, 0
	maxidx1 := -1
	maxidx2, maxidx3 := make([]int, 0), make([]int, 0)
	for i := k * 2; i < n; i++ {
		sum1 += nums[i-k*2]
		sum2 += nums[i-k]
		sum3 += nums[i]
		if i >= 3*k-1 {
			if sum1 > maxsum1 {
				maxsum1 = sum1
				maxidx1 = i - k*3 + 1
			}
			if sum2+maxsum1 > maxsum2 {
				maxsum2 = sum2 + maxsum1
				maxidx2 = []int{maxidx1, i - k*2 + 1}
			}
			if sum3+maxsum2 > maxsum {
				maxsum = sum3 + maxsum2
				maxidx3 = append(maxidx2, i-k+1)
			}
			sum1 -= nums[i-k*3+1]
			sum2 -= nums[i-k*2+1]
			sum3 -= nums[i-k+1]
		}
	}
	return maxidx3
}