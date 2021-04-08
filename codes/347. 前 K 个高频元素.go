// 用快速排序的过程中，根据k进行递降
type NumFrequency struct {
	num       int
	frequency int
}

func quickSort(nums []NumFrequency, begin int, end int, k int) []int {
	pivot := nums[begin]
	nums[begin], nums[end] = nums[end], nums[begin]
	i := 0
	j := 0
	for j != end {
		if nums[j].frequency <= pivot.frequency {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			j++
		}
	}
	if j-i+1 <= k {
		result := make([]int, j-i+1)
		for m := 0; m < j-i+1; m++ {
			result[m] = nums[i+m].num
		}
		if j-i+1 < k {
			result = append(result, quickSort(nums, begin, i-1, k-(j-i+1))...)
		}
		return result
	} else {
		return quickSort(nums, i, end, k)
	}
}

func topKFrequent(nums []int, k int) []int {
	numCnt := make(map[int]int)
	for _, num := range nums {
		numCnt[num]++
	}
	numFreqs := make([]NumFrequency, 0)
	for key, val := range numCnt {
		numFreqs = append(numFreqs, NumFrequency{num: key, frequency: val})
	}
	return quickSort(numFreqs, 0, len(numFreqs)-1, k)
}