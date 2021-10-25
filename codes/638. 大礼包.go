// 递归
func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func addAry(a []int, b []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		a[i] += b[i]
	}
}

func subAry(a []int, b []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		a[i] -= b[i]
	}
}

func noLessThan(a []int, b []int) bool {
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] < b[i] {
			return false
		}
	}
	return true
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)
	if len(special) == 0 {
		deal := 0
		for i := 0; i < n; i++ {
			deal += price[i] * needs[i]
		}
		return deal
	}
	bought := make([]int, n)
	deal := math.MaxInt32
	for i := 0; noLessThan(needs, bought); i++ {
		subAry(needs, bought)
		deal = min(deal, shoppingOffers(price, special[1:], needs)+special[0][n]*i)
		addAry(needs, bought)
		addAry(bought, special[0][:n])
	}
	return deal
}