func minimumCardPickup(cards []int) int {
	lastCard := make(map[int]int)
	mindis := math.MaxInt32
	find := false
	for i, card := range cards {
		if last, ok := lastCard[card]; ok && i-last < mindis {
			mindis = i - last
			find = true
		}
		lastCard[card] = i
	}
	if find {
		return mindis + 1
	}
	return -1
}