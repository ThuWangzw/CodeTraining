func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findKDistantIndices(nums []int, key int, k int) []int {
	ans := make([]int, 0)
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == key {
			begin := max(0, i-k)
			end := min(n-1, i+k)
			if len(ans) > 0 {
				begin = max(begin, ans[len(ans)-1]+1)
			}
			for j := begin; j <= end; j++ {
				ans = append(ans, j)
			}
		}
	}
	return ans
}