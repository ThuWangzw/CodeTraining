func reverseOnlyLetters(s string) string {
	sbs := []byte(s)
	i := 0
	j := len(sbs) - 1
	for i < j {
		for (s[i] < 'a' || s[i] > 'z') && (s[i] < 'A' || s[i] > 'Z') && i < j {
			i++
		}
		for (s[j] < 'a' || s[j] > 'z') && (s[j] < 'A' || s[j] > 'Z') && i < j {
			j--
		}
		sbs[i], sbs[j] = sbs[j], sbs[i]
		i++
		j--
	}
	return string(sbs)
}