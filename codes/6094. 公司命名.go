func distinctNames(ideas []string) int64 {
	dict := make(map[string]map[byte]bool)
	cdict := make(map[byte]map[byte]int) // cdict[a][b]表示首字母为a但是不能为b的数量
	for c := 'a'; c <= 'z'; c++ {
		cdict[byte(c)] = make(map[byte]int)
	}
	for _, idea := range ideas {
		half := idea[1:]
		if _, ok := dict[half]; !ok {
			dict[half] = make(map[byte]bool)
		}
		dict[half][idea[0]] = true
	}
	for _, idea := range ideas {
		half := idea[1:]
		for c := 'a'; c <= 'z'; c++ {
			if !dict[half][byte(c)] {
				cdict[idea[0]][byte(c)]++
			}

		}
	}
	var cnt int64
	for _, idea := range ideas {
		half := idea[1:]
		for c := 'a'; c <= 'z'; c++ {
			if !dict[half][byte(c)] {
				cnt += int64(cdict[byte(c)][idea[0]])
			}
		}
	}
	return cnt
}