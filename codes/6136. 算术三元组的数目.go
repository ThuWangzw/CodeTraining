func arithmeticTriplets(nums []int, diff int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	ans := 0
	for num, cnt := range numMap {
		ans += cnt * numMap[num+diff] * numMap[num+diff*2]
	}
	return ans
}