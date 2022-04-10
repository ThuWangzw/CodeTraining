func sortArray(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	mid := n / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])
	sortedNums := make([]int, n)
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sortedNums[i+j] = left[i]
			i++
		} else {
			sortedNums[i+j] = right[j]
			j++
		}
	}
	for i < len(left) {
		sortedNums[i+j] = left[i]
		i++
	}
	for j < len(right) {
		sortedNums[i+j] = right[j]
		j++
	}
	return sortedNums
}

func maximumProduct(nums []int, k int) int {
	n := len(nums)
	nums = sortArray(nums)
	last := nums[0]
	i := 1
	for ; i < n; i++ {
		if k >= (nums[i]-last)*i {
			k -= (nums[i] - last) * i
			last = nums[i]
		} else {
			break
		}
	}
	last1 := last + k/i
	last2 := last + k/i + 1
	last2cnt := k % i
	last1cnt := i - (k % i)
	mod := 1000000007
	ans := 1
	for j := 0; j < last1cnt; j++ {
		ans = (ans * last1) % mod
	}
	for j := 0; j < last2cnt; j++ {
		ans = (ans * last2) % mod
	}
	for ; i < n; i++ {
		ans = (ans * nums[i]) % mod
	}
	return ans
}