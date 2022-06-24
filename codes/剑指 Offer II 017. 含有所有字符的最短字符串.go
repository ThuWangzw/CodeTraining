func minWindow(s string, t string) string {
	tmap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tmap[t[i]]++
	}
	smap := make(map[byte]int)
	diff := len(tmap)
	left := 0
	ans := ""
	for right := 0; right < len(s); right++ {
		c := s[right]
		smap[c]++
		if smap[c] == tmap[c] {
			diff--
		}
		if diff == 0 {
			for left <= right {
				if len(ans) == 0 || right-left+1 < len(ans) {
					ans = s[left : right+1]
				}
				smap[s[left]]--
				if smap[s[left]] < tmap[s[left]] {
					diff++
					left++
					break
				}
				left++
			}
		}
	}
	return ans
}