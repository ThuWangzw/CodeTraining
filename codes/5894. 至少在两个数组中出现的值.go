func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	appear := make([]map[int]bool, 101)
	for i := 0; i <= 100; i++ {
		appear[i] = make(map[int]bool)
	}
	for _, num := range nums1 {
		appear[num][1] = true
	}
	for _, num := range nums2 {
		appear[num][2] = true
	}
	for _, num := range nums3 {
		appear[num][3] = true
	}
	ans := make([]int, 0)
	for num, cnts := range appear {
		if len(cnts) > 1 {
			ans = append(ans, num)
		}
	}
	return ans
}