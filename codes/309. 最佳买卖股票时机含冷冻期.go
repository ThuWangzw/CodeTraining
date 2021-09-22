// 动态规划，方法一：根据股票问题的通解来做
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int) int {
	n := len(prices)
	k := n / 2
	buy := make([][]int, n+1)
	buy[0] = make([]int, k+1)
	for i := 0; i <= k; i++ {
		buy[0][i] = math.MinInt32
	}
	sell := make([][]int, n+1)
	sell[0] = make([]int, k+1)
	for i := 1; i <= n; i++ {
		buy[i] = make([]int, k+1)
		buy[i][0] = math.MinInt32
		sell[i] = make([]int, k+1)
		for j := 1; j <= k; j++ {
			sell[i][j] = max(buy[i-1][j]+prices[i-1], sell[i-1][j])
			if i > 1 {
				buy[i][j] = max(sell[i-2][j-1]-prices[i-1], buy[i-1][j])
			} else {
				buy[i][j] = max(-prices[i-1], buy[i-1][j])
			}
		}
	}
	maxPro := 0
	for i := 0; i <= k; i++ {
		maxPro = max(maxPro, sell[n][i])
	}
	return maxPro
}

// 方法二：另寻一个新状态
func maxProfit(prices []int) int {
	n := len(prices)
	ans := []int{-prices[0], 0, 0}
	maxPro := 0
	for i := 1; i < n; i++ {
		oldans := []int{ans[0], ans[1], ans[2]}
		ans[0] = max(oldans[0], oldans[2]-prices[i])
		ans[1] = oldans[0] + prices[i]
		ans[2] = max(oldans[2], oldans[1])
		maxPro = max(ans[1], ans[2])
	}
	return maxPro
}