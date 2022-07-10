func fillCups(amount []int) int {
	cnt := 0
	for {
		sort.Ints(amount)
		if amount[2] == 0 {
			return cnt
		}
		amount[2]--
		if amount[1] > 0 {
			amount[1]--
		}
		cnt++
	}
}