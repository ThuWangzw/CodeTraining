func missingNumber(nums []int) int {
	n := len(nums)
	sum := (1 + n) * n / 2
	for _, num := range nums {
		sum -= num
	}
	return sum
}