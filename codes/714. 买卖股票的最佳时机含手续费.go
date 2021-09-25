// 股票问题
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	buy := make([]int, n+1)
	sell := make([]int, n+1)
	buy[0] = math.MinInt32
	for i := 1; i <= n; i++ {
		buy[i] = max(sell[i-1]-prices[i-1]-fee, buy[i-1])
		sell[i] = max(buy[i-1]+prices[i-1], sell[i-1])
	}
	return sell[n]
}