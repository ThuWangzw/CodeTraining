func qsort(arr []int, begin int, end int) {
	if begin >= end {
		return
	}
	pivot := arr[begin]
	left := begin + 1
	right := end
	for left <= right {
		if arr[left] <= pivot {
			left++
		} else {
			arr[left], arr[right] = arr[right], arr[left]
			right--
		}
	}
	arr[begin], arr[right] = arr[right], arr[begin]
	qsort(arr, begin, right-1)
	qsort(arr, right+1, end)
}

func canReorderDoubled(arr []int) bool {
	n := len(arr)
	qsort(arr, 0, n-1)
	numMap := make(map[int]int)
	for _, num := range arr {
		numMap[num]++
	}
	for _, num := range arr {
		if numMap[num] != 0 {
			numMap[num]--
			if num < 0 {
				if num%2 != 0 || numMap[num/2] == 0 {
					return false
				} else {
					numMap[num/2]--
				}
			} else {
				if numMap[num*2] == 0 {
					return false
				} else {
					numMap[num*2]--
				}
			}
		}
	}
	return true
}