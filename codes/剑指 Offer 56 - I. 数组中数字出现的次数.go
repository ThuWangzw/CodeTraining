func singleNumbers(nums []int) []int {
	c := 0
	for _, num := range nums {
		c ^= num
	}
	flag := 1
	for flag&c == 0 {
		flag = flag << 1
	}
	a := 0
	b := 0
	for _, num := range nums {
		if num&flag == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}