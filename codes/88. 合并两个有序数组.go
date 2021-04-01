// 太强了吧，从后向前归并数组，可以避免额外的空间复杂度
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[i+j+1] = nums1[i]
			i--
		} else if nums1[i] < nums2[j] {
			nums1[i+j+1] = nums2[j]
			j--
		}
	}
	for j >= 0 {
		nums1[i+j+1] = nums2[j]
		j--
	}
}