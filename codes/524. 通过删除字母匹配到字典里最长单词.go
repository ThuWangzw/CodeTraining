// 简单的双指针直接比较题
func containWord(s string, t string) bool {
	m := len(s)
	n := len(t)
	i := 0
	j := 0
	for i != m && j != n {
		if s[i] == t[j] {
			j++
		}
		i++
	}
	return j == n
}

func LargeThan(s string, t string) bool {
	m := len(s)
	n := len(t)
	if m != n {
		return m > n
	} else if m == 0 {
		return false
	} else {
		return s < t
	}
}

func findLongestWord(s string, dictionary []string) string {
	max_str := ""
	for _, t := range dictionary {
		if LargeThan(t, max_str) && containWord(s, t) {
			max_str = t
		}
	}
	return max_str
}