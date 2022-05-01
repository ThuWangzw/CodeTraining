func minArray(numbers []int) int {
	begin := 0
	end := len(numbers) - 1
	for end-begin > 3 {
		if numbers[end] > numbers[begin] {
			return numbers[begin]
		}
		mid := (begin + end) / 2
		if numbers[mid] < numbers[begin] {
			end = mid
		} else if numbers[mid] > numbers[begin] {
			begin = mid
		} else {
			begin++
		}
	}
	minv := numbers[begin]
	for i := begin + 1; i <= end; i++ {
		if minv > numbers[i] {
			minv = numbers[i]
		}
	}
	return minv
}