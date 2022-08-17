var dp [][][][]int
var nums []int

func isLimit(limit, num, cur int) int {
	if limit == 0 {
		return 0
	}
	if num == cur {
		return 1
	}
	return 0
}

func dfs(pos, mask, limit, zero int) int {
	if pos < 0 {
		return 1 - zero
	}
	if dp[pos][mask][limit][zero] != -1 {
		return dp[pos][mask][limit][zero]
	}
	dp[pos][mask][limit][zero] = 0

	maxi := 9
	if limit == 1 {
		maxi = nums[pos]
	}
	if zero == 1 {
		dp[pos][mask][limit][zero] = dfs(pos-1, mask, 0, 1)
		for i := 1; i <= maxi; i++ {
			if mask&(1<<i) == 0 {
				dp[pos][mask][limit][zero] += dfs(pos-1, mask|(1<<i), isLimit(limit, nums[pos], i), 0)
			}
		}
	} else {
		for i := 0; i <= maxi; i++ {
			if mask&(1<<i) == 0 {
				dp[pos][mask][limit][zero] += dfs(pos-1, mask|(1<<i), isLimit(limit, nums[pos], i), 0)
			}
		}
	}
	return dp[pos][mask][limit][zero]
}

func countSpecialNumbers(n int) int {
	dp = make([][][][]int, 11)
	for pos := 0; pos < 11; pos++ {
		dp[pos] = make([][][]int, 1024)
		for mask := 0; mask < 1024; mask++ {
			dp[pos][mask] = make([][]int, 2)
			for limit := 0; limit < 2; limit++ {
				dp[pos][mask][limit] = []int{-1, -1}
			}
		}
	}

	nums = make([]int, 0)
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
	}
	return dfs(len(nums)-1, 0, 1, 1)
}