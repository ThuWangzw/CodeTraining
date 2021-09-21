func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	maxPro := 0
	for _, price := range prices {
		maxPro = max(maxPro, price-minPrice)
		minPrice = min(minPrice, price)
	}
	return maxPro
}