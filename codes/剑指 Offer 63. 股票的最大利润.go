func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit := 0
	lastmin := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-lastmin > profit {
			profit = prices[i] - lastmin
		}
		if prices[i] < lastmin {
			lastmin = prices[i]
		}
	}
	return profit
}