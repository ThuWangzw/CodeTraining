func canPlaceFlowers(flowerbed []int, n int) bool {
	maxFlower := 0
	last := -2
	m := len(flowerbed) + 2
	var end = []int{0, 1}
	flowerbed = append(flowerbed, end...)
	for i := 0; i < m; i++ {
		if flowerbed[i] == 1 {
			if i-last > 1 {
				maxFlower += (i - last - 2) / 2
			}
			last = i
		}
	}
	return n <= maxFlower
}