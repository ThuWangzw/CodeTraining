func dicesProbability(n int) []float64 {
	last := []float64{float64(1) / float64(6), float64(1) / float64(6), float64(1) / float64(6), float64(1) / float64(6), float64(1) / float64(6), float64(1) / float64(6)}
	for k := 2; k <= n; k++ {
		current := make([]float64, k*6)
		for i := 0; i < len(last); i++ {
			for j := 1; j <= 6; j++ {
				current[i+j] += last[i] / 6
			}
		}
		last = current
	}
	return last[n-1:]
}