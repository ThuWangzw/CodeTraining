// 经典单调栈
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make(map[int]int)
	m := len(nums1)
	n := len(nums2)
	for _, num := range nums1 {
		ans[num] = -1
	}
	stack := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if _, ok := ans[nums2[i]]; ok && len(stack) > 0 {
			ans[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	ansary := make([]int, m)
	for i := 0; i < m; i++ {
		ansary[i] = ans[nums1[i]]
	}
	return ansary
}