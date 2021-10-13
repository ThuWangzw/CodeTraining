// 求欧拉回路
func visit(numRemains [][]int, cur int, n int, k int, maxPass int, ans *[]byte) {
	for len(numRemains[cur]) > 0 {
		edge := numRemains[cur][0]
		numRemains[cur] = numRemains[cur][1:]
		next := (cur*k + edge) % maxPass
		visit(numRemains, next, n, k, maxPass, ans)
	}
	*ans = append([]byte{byte(cur%k + '0')}, *ans...)
}

func crackSafe(n int, k int) string {
	ans := []byte{}
	if n == 1 {
		for i := 0; i < k; i++ {
			ans = append(ans, byte(i)+'0')
		}
		return string(ans)
	}
	if k == 1 {
		for i := 0; i < n; i++ {
			ans = append(ans, '0')
		}
		return string(ans)
	}
	maxPass := int(math.Pow(float64(k), float64(n-1)))
	numRemains := make([][]int, maxPass)
	for i := 0; i < maxPass; i++ {
		numRemains[i] = make([]int, k)
		for j := 0; j < k; j++ {
			numRemains[i][j] = j
		}
	}
	visit(numRemains, 0, n, k, maxPass, &ans)
	for i := 0; i < n-2; i++ {
		ans = append([]byte{'0'}, ans...)
	}
	return string(ans)
}