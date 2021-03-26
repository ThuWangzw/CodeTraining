// 贪心算法
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	min_rates := make([]int, 0)
	for i := 0; i < n; i++ {
		if (i == 0 || ratings[i] <= ratings[i-1]) && (i == n-1 || ratings[i] <= ratings[i+1]) {
			min_rates = append(min_rates, i)
		}
	}
	for _, min_rate := range min_rates {
		candies[min_rate] = 1
		for i := min_rate - 1; i >= 0 && ratings[i] > ratings[i+1]; i-- {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
		for i := min_rate + 1; i < n && ratings[i] > ratings[i-1]; i++ {
			candies[i] = max(candies[i], candies[i-1]+1)
		}
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += candies[i]
	}
	return sum
}