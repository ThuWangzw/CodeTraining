func mySqrt(x int) int {
	begin := 0
	end := x
	for begin <= end {
		mid := (begin + end) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return end
}