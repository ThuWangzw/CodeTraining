func isValid(nums []int) bool {
	return len(nums) > 0 && nums[len(nums)-1] != 0
}

func isPower(nums []int) bool {
	num := 0
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		num = num*10 + nums[i]
	}
	for num > 1 {
		if num%2 > 0 {
			return false
		}
		num = num >> 1
	}
	return true
}

func getReorder(nums []int, idx int, cur []int) bool {
	n := len(nums)
	if idx == n {
		return isValid(cur) && isPower(cur)
	}
	num := nums[idx]
	for i := 0; i < n; i++ {
		if cur[i] == -1 {
			cur[i] = num
			if flag := getReorder(nums, idx+1, cur); flag {
				return true
			}
			cur[i] = -1
		}
	}
	return false
}

func reorderedPowerOf2(n int) bool {
	nums := make([]int, 0)
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
	}
	cur := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		cur[i] = -1
	}
	return getReorder(nums, 0, cur)
}