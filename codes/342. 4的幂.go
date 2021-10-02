func isPowerOfFour(n int) bool {
	if (n > 0) && (n&(n-1) == 0) && (n&0x55555555 != 0) {
		return true
	}
	return false
}