func minimumDeletions(s string) int {
	alla := 0
	for _, c := range s {
		if c == 'a' {
			alla++
		}
	}
	nowa := 0
	nowb := 0
	ans := min(alla, len(s)-alla)
	for _, c := range s {
		if c == 'b' {
			ans = min(ans, nowb+alla-nowa)
			nowb++
		} else {
			nowa++
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}