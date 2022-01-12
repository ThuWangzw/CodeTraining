// 有点偏脑筋急转弯
func increasingTriplet(nums []int) bool {
	minnum := math.MaxInt32
	minnum2 := math.MaxInt32
	for _, num := range nums {
		if num > minnum && num > minnum2 {
			return true
		}
		if num <= minnum {
			minnum = num
		} else if num < minnum2 {
			minnum2 = num
		}
	}
	return false
}