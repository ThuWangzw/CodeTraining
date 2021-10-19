// 贪心算法，写两个快排即可
func quickSort(nums []int) {
	num := nums[0]
	n := len(nums)
	left := 1
	right := n - 1
	for left <= right {
		if nums[left] < num {
			left++
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}
	nums[0], nums[right] = nums[right], nums[0]
	if right > 0 {
		quickSort(nums[:right])
	}
	if right+1 < n {
		quickSort(nums[right+1:])
	}
}

func quickSortIndex(nums []int, idxes []int) {
	num := nums[idxes[0]]
	n := len(idxes)
	left := 1
	right := n - 1
	for left <= right {
		if nums[idxes[left]] < num {
			left++
		} else {
			idxes[left], idxes[right] = idxes[right], idxes[left]
			right--
		}
	}
	idxes[0], idxes[right] = idxes[right], idxes[0]
	if right > 0 {
		quickSortIndex(nums, idxes[:right])
	}
	if right+1 < n {
		quickSortIndex(nums, idxes[right+1:])
	}
}

func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	quickSort(nums1)
	idxes := make([]int, n)
	for i := 0; i < n; i++ {
		idxes[i] = i
	}
	quickSortIndex(nums2, idxes)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	left := 0
	right := n - 1
	for _, num := range nums1 {
		if num > nums2[idxes[left]] {
			ans[idxes[left]] = num
			left++
		} else {
			ans[idxes[right]] = num
			right--
		}
	}
	return ans
}