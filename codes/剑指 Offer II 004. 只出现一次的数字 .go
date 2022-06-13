func singleNumber(nums []int) int {
	l := 0
	r := 0
	for _, n := range nums {
		l, r = (r&n)|(l&^r&^n), ^l&(r^n)
	}
	return r
}