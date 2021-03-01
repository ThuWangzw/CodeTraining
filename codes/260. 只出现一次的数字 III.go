// 是异或操作的变体，思路在于将所有数字分成两类
func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor = xor ^ num
	}

	mask := xor & (-xor)

	result := []int{0, 0}
	for _, num := range nums {
		if mask&num == 0 {
			result[0] = result[0] ^ num
		} else {
			result[1] = result[1] ^ num
		}
	}
	return result
}