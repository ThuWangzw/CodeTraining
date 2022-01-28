func large(propertyA []int, propertyB []int) bool {
	return (propertyA[0] > propertyB[0]) || (propertyA[0] == propertyB[0] && propertyA[1] < propertyB[1])
}

func qsort(properties [][]int, begin int, end int) {
	if begin >= end {
		return
	}
	pivot := properties[begin]
	left := begin + 1
	right := end
	for left <= right {
		if large(properties[left], pivot) {
			left++
		} else {
			properties[left], properties[right] = properties[right], properties[left]
			right--
		}
	}
	properties[begin], properties[right] = properties[right], properties[begin]
	qsort(properties, begin, right-1)
	qsort(properties, right+1, end)
}

func numberOfWeakCharacters(properties [][]int) int {
	n := len(properties)
	qsort(properties, 0, n-1)
	maxdefense := 0
	count := 0
	for _, property := range properties {
		if property[1] < maxdefense {
			count++
		} else {
			maxdefense = property[1]
		}
	}
	return count
}