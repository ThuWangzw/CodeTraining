func countAsterisks(s string) int {
	acnt := 0
	scnt := 0
	for _, c := range s {
		if c == '|' {
			acnt++
		} else if c == '*' && acnt%2 == 0 {
			scnt++
		}
	}
	return scnt
}