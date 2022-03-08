// 预处理即可，极限情况分析一下
func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	leftCandle := make([]int, n)
	rightCandle := make([]int, n)
	candleIndex := make(map[int]int)
	left := -1
	candleCount := 0
	for i := 0; i < n; i++ {
		if s[i] == '|' {
			left = i
			candleIndex[i] = candleCount
			candleCount++
		}
		leftCandle[i] = left
	}
	right := n
	for i := n - 1; i >= 0; i-- {
		if s[i] == '|' {
			right = i
		}
		rightCandle[i] = right
	}
	m := len(queries)
	ans := make([]int, m)
	for i, query := range queries {
		newleft := rightCandle[query[0]]
		newright := leftCandle[query[1]]
		if newleft >= newright {
			ans[i] = 0
		} else {
			ans[i] = newright - newleft - (candleIndex[newright] - candleIndex[newleft])
		}
	}
	return ans
}