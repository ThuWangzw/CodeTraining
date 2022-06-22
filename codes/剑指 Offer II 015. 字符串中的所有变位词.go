func findAnagrams(s string, p string) []int {
	m := len(s)
	n := len(p)
	if m < n {
		return []int{}
	}
	pmap := make(map[rune]int)
	for _, c := range p {
		pmap[c]++
	}
	diff := 0
	smap := make(map[rune]int)
	for i := 0; i < n; i++ {
		smap[rune(s[i])]++
	}
	for i := 'a'; i <= 'z'; i++ {
		if pmap[i] != smap[i] {
			diff++
		}
	}
	ans := make([]int, 0)
	if diff == 0 {
		ans = append(ans, 0)
	}
	for i := n; i < m; i++ {
		addc := rune(s[i])
		deletec := rune(s[i-n])
		smap[addc]++
		if smap[addc] == pmap[addc] {
			diff--
		} else if smap[addc]-1 == pmap[addc] {
			diff++
		}
		smap[deletec]--
		if smap[deletec] == pmap[deletec] {
			diff--
		} else if smap[deletec]+1 == pmap[deletec] {
			diff++
		}
		if diff == 0 {
			ans = append(ans, i-n+1)
		}
	}
	return ans
}