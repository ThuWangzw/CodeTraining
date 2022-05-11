func quickSelect(arr []int, k int, begin int, end int) []int {
	if begin > end || k == 0 {
		return []int{}
	}
	pivot := arr[begin]
	left, right := begin, end
	for left <= right {
		if arr[left] <= pivot {
			left++
		} else {
			arr[left], arr[right] = arr[right], arr[left]
			right--
		}
	}
	arr[begin], arr[right] = arr[right], arr[begin]
	if right-begin+1 == k {
		return arr[begin : right+1]
	} else if right-begin+1 > k {
		return quickSelect(arr, k, begin, right-1)
	} else {
		return append(arr[begin:right+1], quickSelect(arr, k-(right-begin+1), right+1, end)...)
	}
}

func getLeastNumbers(arr []int, k int) []int {
	return quickSelect(arr, k, 0, len(arr)-1)
}