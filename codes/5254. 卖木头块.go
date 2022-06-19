func sellingWood(m int, n int, prices [][]int) int64 {
	maxprices := make([][]int64, m+1)
	for i := 0; i <= m; i++ {
		maxprices[i] = make([]int64, n+1)
	}
	for _, price := range prices {
		maxprices[price[0]][price[1]] = int64(price[2])
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			maxprices[i][j] = max(max(maxprices[i-1][j], maxprices[i][j-1]), maxprices[i][j])
		}
	}
	ans := make([][]int64, m+1)
	for i := 0; i <= m; i++ {
		ans[i] = make([]int64, n+1)
		for j := 0; j <= n; j++ {
			ans[i][j] = -1
		}
	}
	return sellingWoodIter(m, n, maxprices, ans)
}

func sellingWoodIter(m int, n int, prices [][]int64, ans [][]int64) int64 {
	if ans[m][n] >= 0 {
		return ans[m][n]
	}
	var maxprice int64 = prices[m][n]
	for i := 1; i < m; i++ {
		maxprice = max(maxprice, sellingWoodIter(i, n, prices, ans)+sellingWoodIter(m-i, n, prices, ans))
	}
	for j := 1; j < n; j++ {
		maxprice = max(maxprice, sellingWoodIter(m, j, prices, ans)+sellingWoodIter(m, n-j, prices, ans))
	}
	ans[m][n] = maxprice
	return maxprice
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}