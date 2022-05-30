func parseMoney(word string) (float64, bool) {
	if word[0] != '$' {
		return -1, false
	}
	money, err := strconv.ParseFloat(word[1:], 64)
	if err == nil && money >= 0 {
		return money, true
	}
	return -1, false
}

func discountPrices(sentence string, discount int) string {
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if money, ok := parseMoney(word); ok {
			money *= 1 - float64(discount)/100
			words[i] = fmt.Sprintf("$%.2f", money)
		}
	}
	return strings.Join(words, " ")
}