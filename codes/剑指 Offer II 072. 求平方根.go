func mySqrt(x int) int {
	begin := 0
	end := x
	for begin < end-3 {
		mid := (begin + end) / 2
		res := int64(mid) * int64(mid)
		if res == int64(x) {
			return mid
		} else if res < int64(x) {
			begin = mid
		} else {
			end = mid - 1
		}
	}
	for i := end; i >= begin; i-- {
		if int64(i)*int64(i) <= int64(x) {
			return i
		}
	}
	return 0
}