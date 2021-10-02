func hammingDistance(x int, y int) int {
	z := x ^ y
	c := 0
	for c = 0; z > 0; c++ {
		z = z & (z - 1)
	}
	return c
}