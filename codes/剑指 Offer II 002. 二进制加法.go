func addBinary(a string, b string) string {
	c := make([]byte, 0)
	m := len(a)
	n := len(b)
	i, j := m-1, n-1
	last := 0
	for i >= 0 || j >= 0 {
		ac := 0
		bc := 0
		if i >= 0 {
			ac = int(a[i] - '0')
			i--
		}
		if j >= 0 {
			bc = int(b[j] - '0')
			j--
		}
		c = append(c, '0'+byte((ac+bc+last)%2))
		last = (ac + bc + last) / 2
	}
	if last > 0 {
		c = append(c, '0'+byte(last))
	}
	k := len(c)
	for i := 0; i < k/2; i++ {
		c[i], c[k-1-i] = c[k-1-i], c[i]
	}
	return string(c)
}