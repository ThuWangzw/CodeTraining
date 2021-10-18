func findComplement(num int) int {
	bit := 0
	newnum := num
	for newnum > 0 {
		newnum = newnum >> 1
		bit++
	}
	mask := 0
	for i := 0; i < bit; i++ {
		mask = (mask << 1) + 1
	}
	return (^num) & mask
}