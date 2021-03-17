// 时间复杂度为总金额的线性背包问题
func min(a int, b int) int {
	if a<b{
		return a
	} else {
		return b
	}
}

func coinChange(coins []int, amount int) int {
	result := make([]int, amount+1, amount+1)
	for i := 1; i <= amount; i++ {
		result[i] = math.MaxInt32
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin>= 0 {
				result[i] = min(result[i], result[i-coin]+1)
			}
		}
	}
	if result[amount] == math.MaxInt32 {
		return -1
	}
	return result[amount]
}