// n和n-1的与，会移除掉n的最后一位1
func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		cnt++
		num = num & (num - 1)
	}
	return cnt
}