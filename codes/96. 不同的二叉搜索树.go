var numTreesMap = make(map[int]int)

func numTrees(n int) int {
	if num, ok := numTreesMap[n]; ok {
		return num
	}
	if n <= 1 {
		numTreesMap[n] = 1
		return 1
	}
	cnt := 0
	for i := 1; i <= n; i++ {
		cnt += numTrees(i-1) * numTrees(n-i)
	}
	numTreesMap[n] = cnt
	return cnt
}