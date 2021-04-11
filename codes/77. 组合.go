var results [][]int

func search(n int, k int, begin int, cur []int) {
	end := n - (k - len(cur)) + 1
	if k == len(cur) {
		mcur := make([]int, k)
		copy(mcur, cur)
		results = append(results, mcur)
	} else {
		for i := begin; i <= end; i++ {
			search(n, k, i+1, append(cur, i))
		}
	}
}

func combine(n int, k int) [][]int {
	results = make([][]int, 0)
	search(n, k, 1, []int{})
	return results
}