func isSelfDividingNumbers(num int) bool {
	if num == 0 {
		return false
	}
	numdivid := num
	for numdivid > 0 {
		if numdivid%10 == 0 || num%(numdivid%10) > 0 {
			return false
		}
		numdivid /= 10
	}
	return true
}

func selfDividingNumbers(left int, right int) []int {
	ans := make([]int, 0)
	for left <= right {
		if isSelfDividingNumbers(left) {
			ans = append(ans, left)
		}
		left++
	}
	return ans
}