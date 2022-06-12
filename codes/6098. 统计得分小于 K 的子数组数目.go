func bs(target int, sums []int64, k int64) int64 {
	begin := 0
	end := target
	for begin < end-3 {
		mid := (begin + end) / 2
		cal := sums[target] * int64(target+1)
		if mid > 0 {
			cal = (sums[target] - sums[mid-1]) * int64(target-mid+1)
		}
		if cal >= k {
			begin = mid + 1
		} else {
			end = mid
		}
	}
	for i := begin; i <= end; i++ {
		cal := sums[target] * int64(target+1)
		if i > 0 {
			cal = (sums[target] - sums[i-1]) * int64(target-i+1)
		}
		if cal < k {
			return int64(target - i + 1)
		}
	}
	return 0
}

func countSubarrays(nums []int, k int64) int64 {
	n := len(nums)
	sums := make([]int64, n)
	sums[0] = int64(nums[0])
	for i := 1; i < n; i++ {
		sums[i] = sums[i-1] + int64(nums[i])
	}
	var ans int64
	for i := 0; i < n; i++ {
		ans += bs(i, sums, k)
	}
	return ans
}