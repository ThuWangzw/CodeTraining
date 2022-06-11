var resultString []string = make([]string, 0)

func permutationMap(dict map[byte]int, candidate []byte) {
	find := false
	for c, cnt := range dict {
		if cnt > 0 {
			find = true
			dict[c]--
			candidate = append(candidate, c)
			permutationMap(dict, candidate)
			dict[c]++
			candidate = candidate[:len(candidate)-1]
		}
	}
	if !find {
		resultString = append(resultString, string(candidate))
	}
}

func permutation(s string) []string {
	resultString = make([]string, 0)
	dict := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}
	candidate := make([]byte, 0)
	permutationMap(dict, candidate)
	return resultString
}