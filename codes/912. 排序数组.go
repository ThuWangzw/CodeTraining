func mergeSort(nums []int, begin int, end int) {
	if end == begin {
		return
	}
	mid := (begin + end) / 2
	mergeSort(nums, begin, mid)
	mergeSort(nums, mid+1, end)
	tmp := make([]int, end-mid)
	for i := 0; i < end-mid; i++ {
		tmp[i] = nums[mid+1+i]
	}
	i := mid
	j := end - mid - 1
	for i >= begin && j >= 0 {
		if nums[i] > tmp[j] {
			nums[i+j+1] = nums[i]
			i--
		} else {
			nums[i+j+1] = tmp[j]
			j--
		}
	}
	for j >= 0 {
		nums[begin+j] = tmp[j]
		j--
	}
}

func sortArray(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}