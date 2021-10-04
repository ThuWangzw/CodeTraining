// 脑筋急转弯,主要分清楚什么时候Alice会赢
func stoneGameIX(stones []int) bool {
	stoneCnt := []int{0, 0, 0}
	for _, stone := range stones {
		stoneCnt[stone%3]++
	}
	if stoneCnt[1] > 0 {
		maxStep := 0
		if stoneCnt[2] <= stoneCnt[1]-2 {
			maxStep = stoneCnt[2]*2 + 2 + stoneCnt[0]
		} else {
			maxStep = stoneCnt[1]*2 - 1 + stoneCnt[0]
		}
		if maxStep%2 == 1 && maxStep != len(stones) {
			return true
		}
	}
	if stoneCnt[2] > 0 {
		maxStep := 0
		if stoneCnt[1] <= stoneCnt[2]-2 {
			maxStep = stoneCnt[1]*2 + 2 + stoneCnt[0]
		} else {
			maxStep = stoneCnt[2]*2 - 1 + stoneCnt[0]
		}
		if maxStep%2 == 1 && maxStep != len(stones) {
			return true
		}
	}
	return false
}