func zeroFilledSubarray(nums []int) int64 {
	var cnt, zeros int64
	for _, num := range nums {
		if num != 0 {
			cnt += (zeros + 1) * zeros / 2
			zeros = 0
		} else {
			zeros++
		}
	}
	cnt += (zeros + 1) * zeros / 2
	return cnt
}