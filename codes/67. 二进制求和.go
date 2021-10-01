func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func addBinary(a string, b string) string {
	res := ""
	alen := len(a)
	blen := len(b)
	maxlen := max(alen, blen)
	c := 0
	for i := 0; i < maxlen; i++ {
		anum := 0
		if alen-1-i >= 0 && a[alen-1-i] == '1' {
			anum = 1
		}
		bnum := 0
		if blen-1-i >= 0 && b[blen-1-i] == '1' {
			bnum = 1
		}
		res = string(byte((anum+bnum+c)%2)+'0') + res
		c = (anum + bnum + c) / 2
	}
	if c == 1 {
		res = "1" + res
	}
	return res
}