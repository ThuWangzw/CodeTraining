// 双指针
func judgeSquareSum(c int) bool {
	a := 0
	b := int(math.Sqrt(float64(c)))
	for a != b {
		res := a*a + b*b - c
		if res == 0 {
			return true
		} else if res < 0 {
			a++
		} else {
			b--
		}
	}
	return (a*a + b*b) == c
}