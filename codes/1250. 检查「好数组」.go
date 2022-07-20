func isGoodArray(nums []int) bool {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans = gcd(ans, nums[i])
	}
	return ans == 1
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}