func findNext(p string) []int {
	n := len(p)
	next := make([]int, n)
	next[0] = -1
	k := -1
	j := 0
	for j < n-1 {
		if k == -1 || p[j] == p[k] {
			j++
			k++
			next[j] = k
		} else {
			k = next[k]
		}
	}
	return next
}

func strStr(haystack string, needle string) int {
	m := len(haystack)
	n := len(needle)
	if n == 0 {
		return 0
	}
	next := findNext(needle)
	j := 0
	i := 0
	for i < m {
		if j == -1 {
			j++
			i++
		} else if haystack[i] == needle[j] {
			j++
			i++
			if j == n {
				return i - n
			}
		} else if haystack[i] != needle[j] {
			j = next[j]
		}
	}
	return -1
}