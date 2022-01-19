func containsNearbyDuplicate(nums []int, k int) bool {
	numsMap := make(map[int]int)
	for i, num := range nums {
		idx, in := numsMap[num]
		if !in {
			numsMap[num] = idx
		}
		if in && i-idx <= k {
			return true
		}
		numsMap[num] = i
	}
	return false
}