// 前缀和map
func subarraySum(nums []int, k int) int {
	sumMap := make(map[int]int)
	sum := 0
	cnts := 0
	sumMap[0] = 1
	for _, num := range nums {
		sum += num
		cnts += sumMap[sum-k]
		sumMap[sum]++
	}
	return cnts
}