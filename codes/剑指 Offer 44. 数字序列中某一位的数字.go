var kbitCount map[int]int = make(map[int]int)

func findNthDigit(n int) int {
	k := 1
	for true {
		kcount := getKbitCount(k)
		if n < kcount*k {
			num := n/k + int(math.Pow(10, float64(k-1)))
			if k == 1 {
				num--
			}
			n = k - 1 - n%k
			for i := 0; i < n; i++ {
				num /= 10
			}
			return num % 10
		} else {
			n -= kcount * k
		}
		k++
	}
	return -1
}

func getKbitCount(k int) int {
	if k == 1 {
		return 10
	}
	if kbitCount[k] > 0 {
		return kbitCount[k]
	}
	kbitCount[k] = 9 * int(math.Pow(10, float64(k-1)))
	return kbitCount[k]
}