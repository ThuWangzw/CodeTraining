// 单调栈经典问题+条件
func smallestSubsequence(s string, k int, letter byte, repetition int) string {
	n := len(s)
	ms := make([]byte, 0, k)
	removed := 0
	remain := 0
	contain := 0
	for _, ch := range s {
		if letter == byte(ch) {
			remain++
		}
	}
	for i := 0; i < n; i++ {
		ch := s[i]
		if len(ms) > 0 {
			top := ms[len(ms)-1]
			for removed < n-k && top > ch && (remain > repetition-contain || top != letter) {
				ms = ms[:len(ms)-1]
				removed++
				if top == letter {
					contain--
				}
				if len(ms) == 0 {
					break
				}
				top = ms[len(ms)-1]
			}
		}
		ms = append(ms, ch)
		if ch == letter {
			remain--
			contain++
		}
	}
	if removed < n-k {
		for i := len(ms) - 1; i >= 0 && removed < n-k; i-- {
			if ms[i] != letter || contain > repetition {
				if ms[i] == letter {
					contain--
				}
				ms = append(ms[:i], ms[i+1:]...)
				removed++
			}
		}
	}
	return string(ms)
}