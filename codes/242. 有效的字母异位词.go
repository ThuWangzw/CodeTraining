func isAnagram(s string, t string) bool {
	scnt := make(map[byte]int)
	tcnt := make(map[byte]int)
	for _, ch := range s {
		scnt[byte(ch)]++
	}
	for _, ch := range t {
		tcnt[byte(ch)]++
	}
	for ch, cnt := range scnt {
		if cnt != tcnt[ch] {
			return false
		}
	}
	for ch, cnt := range tcnt {
		if cnt != scnt[ch] {
			return false
		}
	}
	return true
}