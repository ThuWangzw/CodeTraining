func bestHand(ranks []int, suits []byte) string {
	suitMap := make(map[byte][]int)
	rankMap := make(map[int][]byte)
	for i := 0; i < 5; i++ {
		if _, ok := suitMap[suits[i]]; !ok {
			suitMap[suits[i]] = make([]int, 0)
		}
		if _, ok := rankMap[ranks[i]]; !ok {
			rankMap[ranks[i]] = make([]byte, 0)
		}
		suitMap[suits[i]] = append(suitMap[suits[i]], ranks[i])
		rankMap[ranks[i]] = append(rankMap[ranks[i]], suits[i])
	}
	if len(suitMap) == 1 {
		return "Flush"
	}
	maxRank := 1
	for _, suits := range rankMap {
		if len(suits) > maxRank {
			maxRank = len(suits)
		}
	}
	if maxRank >= 3 {
		return "Three of a Kind"
	}
	if maxRank == 2 {
		return "Pair"
	}
	return "High Card"
}