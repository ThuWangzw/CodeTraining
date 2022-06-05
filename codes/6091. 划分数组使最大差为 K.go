func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	begin := 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]-nums[begin] > k {
			ans++
			begin = i
		}
	}
	ans++
	return ans
}