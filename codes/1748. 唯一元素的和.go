func sumOfUnique(nums []int) int {
	numFlags := make(map[int]int) // 0, 1, 2 = visit 0, 1, >=2
	ans := 0
	for _, num := range nums {
		switch numFlags[num] {
		case 0:
			numFlags[num] = 1
			ans += num
		case 1:
			numFlags[num] = 2
			ans -= num
		}
	}
	return ans
}