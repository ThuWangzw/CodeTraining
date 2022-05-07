func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func findK(nums []int, k int) int {
	return nums[k-1]
}

func findKSortedArray(nums1 []int, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return findK(nums2, k)
	}
	if len(nums2) == 0 {
		return findK(nums1, k)
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}
	idx := k/2 - 1
	aidx, bidx := idx, idx
	if len(nums1) <= aidx {
		aidx = len(nums1) - 1
	}
	if len(nums2) <= bidx {
		bidx = len(nums2) - 1
	}
	if nums1[aidx] < nums2[bidx] {
		return findKSortedArray(nums1[aidx+1:], nums2, k-aidx-1)
	} else {
		return findKSortedArray(nums1, nums2[bidx+1:], k-bidx-1)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if (m+n)%2 == 1 {
		return float64(findKSortedArray(nums1, nums2, (m+n+1)/2))
	} else {
		return float64(findKSortedArray(nums1, nums2, (m+n)/2)+findKSortedArray(nums1, nums2, (m+n)/2+1)) / 2
	}
}