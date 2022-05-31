func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	cnt3 := n / 3
	cnt2 := n % 3
	if cnt2 == 0 {
		return int(math.Pow(3, float64(cnt3)))
	} else if cnt2 == 1 {
		return int(math.Pow(3, float64(cnt3-1))) * 4
	} else {
		return int(math.Pow(3, float64(cnt3))) * 2
	}
}