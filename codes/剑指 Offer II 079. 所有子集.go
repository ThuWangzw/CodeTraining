var ans [][]int

func copyAry(a []int) []int {
	b := make([]int, 0, len(a))
	for _, i := range a {
		b = append(b, i)
	}
	return b
}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, make([]int, 0))
	for _, top := range nums {
		n := len(ans)
		for i := 0; i < n; i++ {
			ans = append(ans, copyAry(append(ans[i], top)))
		}
	}
	return ans
}