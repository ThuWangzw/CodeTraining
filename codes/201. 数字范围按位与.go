// 根据按位与的特性不难求解

func rangeBitwiseAnd(m int, n int) int {
	largemask := math.MaxInt32
	bitmask := 1
	result := m
	for i := 0; i < 32; i++ {
		if (m>>i)&1 == 1 {
			tmp := (m + bitmask) & largemask
			if (tmp >= m) && (tmp <= n) {
				result &= ^bitmask
			}
		}
		bitmask = bitmask << 1
		largemask = largemask << 1
	}
	return result
}