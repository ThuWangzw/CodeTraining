// 哈希表简单题
func containsDuplicate(nums []int) bool {
	contain := make(map[int]bool)
	for _, num := range nums {
		if contain[num] {
			return true
		}
		contain[num] = true
	}
	return false
}