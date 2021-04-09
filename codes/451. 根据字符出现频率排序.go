type NumFrequency struct {
	num       byte
	frequency int
}

func mergeSort(numFreqs []NumFrequency) {
	n := len(numFreqs)
	if n < 2 {
		return
	}
	mid := (n + 1) / 2
	mergeSort(numFreqs[:mid])
	mergeSort(numFreqs[mid:])
	// merge
	half := make([]NumFrequency, n-mid)
	copy(half, numFreqs[mid:])
	i := mid - 1
	j := n - mid - 1
	for (i >= 0) && (j >= 0) {
		if numFreqs[i].frequency < half[j].frequency {
			numFreqs[i+j+1] = numFreqs[i]
			i--
		} else {
			numFreqs[i+j+1] = half[j]
			j--
		}
	}
	for j >= 0 {
		numFreqs[j] = half[j]
		j--
	}
}

func frequencySort(s string) string {
	// init num frequency
	charCnts := make(map[byte]int)
	for _, c := range s {
		charCnts[byte(c)]++
	}
	numFreq := make([]NumFrequency, 0)
	for key, val := range charCnts {
		numFreq = append(numFreq, NumFrequency{num: key, frequency: val})
	}
	// mergesort
	mergeSort(numFreq)
	result := ""
	for _, numfreq := range numFreq {
		result += strings.Repeat(string([]byte{numfreq.num}), numfreq.frequency)
	}
	return result
}