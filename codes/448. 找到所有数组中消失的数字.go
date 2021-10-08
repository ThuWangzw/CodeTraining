// 用数组本身当哈希表
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		num := nums[i]
		if num > n {
			num -= n
		}
		if nums[num-1] <= n {
			nums[num-1] += n
		}
	}
	ans := make([]int, 0)
	for i, num := range nums {
		if num <= n {
			ans = append(ans, i+1)
		}
	}
	return ans
}