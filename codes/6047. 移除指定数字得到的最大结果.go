func removeDigit(number string, digit byte) string {
	idx := 0
	last := -1
	for ; idx < len(number); idx++ {
		if number[idx] == digit {
			last = idx
			if idx < len(number)-1 && number[idx] < number[idx+1] {
				return number[:idx] + number[idx+1:]
			}
		}
	}
	return number[:last] + number[last+1:]
}