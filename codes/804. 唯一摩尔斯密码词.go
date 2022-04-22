func uniqueMorseRepresentations(words []string) int {
	word2encoding := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-",
		".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	encodingMap := make(map[string]struct{})
	for _, word := range words {
		s := ""
		for _, ch := range word {
			s += word2encoding[int(ch-'a')]
		}
		encodingMap[s] = struct{}{}
	}
	return len(encodingMap)
}