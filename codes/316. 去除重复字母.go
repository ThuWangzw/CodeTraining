// 单调栈+限制条件
func removeDuplicateLetters(s string) string {
	chCount := make(map[byte]int)
	ifContain := make(map[byte]bool)
	n := len(s)
	ms := make([]byte, 0)
	for _, ch := range s {
		chCount[byte(ch)]++
	}

	for i := 0; i < n; i++ {
		ch := s[i]
		if ifContain[ch] {
			chCount[ch]--
			continue
		} else {
			if len(ms) > 0 {
				top := ms[len(ms)-1]
				for chCount[top] != 0 && top >= ch {
					ifContain[top] = false
					ms = ms[:len(ms)-1]
					if len(ms) == 0 {
						break
					}
					top = ms[len(ms)-1]
				}
			}
			ms = append(ms, ch)
			ifContain[ch] = true
			chCount[ch]--
		}
	}
	return string(ms)
}