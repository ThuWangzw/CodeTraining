// 二分法，注意终止条件即可
func peakIndexInMountainArray(arr []int) int {
	start := 0
	end := len(arr) - 1
	for start < end {
		if end-start == 1 {
			if arr[end] > arr[start] {
				return end
			} else {
				return start
			}
		}
		mid := (start + end) / 2
		if arr[mid] > arr[mid-1] {
			start = mid
		} else {
			end = mid
		}
	}
	return -1
}