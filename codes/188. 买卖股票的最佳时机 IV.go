// 买卖股票问题，动态规划通解：https://leetcode-cn.com/circle/article/qiAgHn/
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(k int, prices []int) int {
	n := len(prices)
	buy := make([]int, k+1)
	for i := 0; i <= k; i++ {
		buy[i] = math.MinInt32
	}
	sell := make([]int, k+1)
	for i := 1; i <= n; i++ {
		for j := k; j >= 1; j-- {
			sell[j] = max(buy[j]+prices[i-1], sell[j])
			buy[j] = max(sell[j-1]-prices[i-1], buy[j])
		}
	}
	maxPro := 0
	for i := 0; i <= k; i++ {
		maxPro = max(maxPro, sell[i])
	}
	return maxPro
}