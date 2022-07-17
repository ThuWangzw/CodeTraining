func minOperations(nums []int, numsDivide []int) int {
	sort.Ints(nums)
	allgcd := numsDivide[0]
	for i := 1; i < len(numsDivide); i++ {
		allgcd = gcd(allgcd, numsDivide[i])
	}
	for i := 0; i < len(nums); i++ {
		if allgcd%nums[i] == 0 {
			return i
		}
	}
	return -1
}

func gcd(a int, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}