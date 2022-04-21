type AlphaCount map[byte]int

func (ac AlphaCount) String() string {
	s := ""
	for i := 'a'; i <= 'z'; i++ {
		s += string(i) + strconv.Itoa(ac[byte(i)])
	}
	return s
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		var alphaCount AlphaCount = make(AlphaCount)
		for _, c := range str {
			alphaCount[byte(c)]++
		}
		key := alphaCount.String()
		if _, ok := groups[key]; !ok {
			groups[key] = make([]string, 0)
		}
		groups[key] = append(groups[key], str)
	}
	ans := make([][]string, 0, len(groups))
	for _, group := range groups {
		ans = append(ans, group)
	}
	return ans
}