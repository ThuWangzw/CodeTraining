func sortIndex(nums []string, indexs []int, k, trim int) int {
	n := len(indexs)
	for i := 0; i < n; i++ {
		for j := 0; j <= n-i-2; j++ {
			if great(nums[indexs[j]], nums[indexs[j+1]], trim) {
				indexs[j], indexs[j+1] = indexs[j+1], indexs[j]
			}
		}
		if k == n-i {
			return indexs[k-1]
		}
	}
	return -1
}

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	ans := make([]int, 0)
	for _, query := range queries {
		indexs := make([]int, len(nums))
		for i := 0; i < len(nums); i++ {
			indexs[i] = i
		}
		ans = append(ans, sortIndex(nums, indexs, query[0], query[1]))
	}
	return ans
}

func great(a string, b string, trim int) bool {
	// a>b -> true
	n := len(a)
	for i := n - trim; i < n; i++ {
		if a[i] > b[i] {
			return true
		} else if a[i] < b[i] {
			return false
		}
	}
	return false
}