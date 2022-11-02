func distance2(a, b []int) float64 {
	return math.Sqrt(float64((a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])))
}

func bestCoordinate(towers [][]int, radius int) []int {
	maxPower := -1
	maxPowerPos := make([]int, 0)
	for x := 0; x <= 50; x++ {
		for y := 0; y <= 50; y++ {
			power := 0
			for _, tower := range towers {
				if d := distance2([]int{x, y}, tower); d <= float64(radius) {
					power += int(float64(tower[2]) / (1 + d))
				}
			}
			if power > maxPower {
				maxPower, maxPowerPos = power, []int{x, y}
			}
		}
	}
	return maxPowerPos
}