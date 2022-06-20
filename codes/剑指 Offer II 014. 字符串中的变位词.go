func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	cmap := make(map[byte]int)
	cmap2 := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		cmap[s1[i]]++
		cmap2[s2[i]]++
	}
	flag := true
	for c := byte('a'); c <= byte('z'); c++ {
		if cmap[c] != cmap2[c] {
			flag = false
			break
		}
	}
	if flag {
		return true
	}
	m := len(s1)
	n := len(s2)
	for i := m; i < n; i++ {
		cmap2[s2[i]]++
		cmap2[s2[i-m]]--
		if cmap[s2[i-m]] == cmap2[s2[i-m]] && cmap[s2[i]] == cmap2[s2[i]] {
			return true
		}
	}
	return false
}