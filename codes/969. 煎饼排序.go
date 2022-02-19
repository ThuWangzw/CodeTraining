func pancakeSort(arr []int) []int {
	n := len(arr)
	ans := make([]int, 0)
	for i := 0; i < n-1; i++ {
		// find max
		maxnum := math.MinInt32
		maxidx := -1
		for j := 0; j < n-i; j++ {
			if arr[j] > maxnum {
				maxnum, maxidx = arr[j], j
			}
		}
		// pancake
		ans = append(ans, maxidx+1)
		for j := 0; j <= maxidx/2; j++ {
			arr[j], arr[maxidx-j] = arr[maxidx-j], arr[j]
		}
		//pancake
		ans = append(ans, n-i)
		for j := 0; j <= (n-1-i)/2; j++ {
			arr[j], arr[n-1-i-j] = arr[n-1-i-j], arr[j]
		}
	}
	return ans
}