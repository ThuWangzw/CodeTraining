// 求递增序列的和
func maxProfit(prices []int) int {
	win := 0
	n := len(prices)
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			win += prices[i] - prices[i-1]
		}
	}
	return win
}