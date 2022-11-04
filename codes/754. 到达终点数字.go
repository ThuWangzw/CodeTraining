func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	k := int(math.Sqrt(float64(target * 2)))
	if target > k*(k+1)/2 {
		k++
	}
	d := k*(k+1)/2 - target
	if d%2 == 0 {
		return k
	} else if k%2 == 0 {
		return k + 1
	} else {
		return k + 2
	}
}