func convert(s string, numRows int) string {
	ans := strings.Builder{}
	base := 2*numRows - 2
	n := len(s)
	if numRows == 1 {
		return s
	}
	for i := 0; i < numRows; i++ {
		k := 0
		for k*base+i < n {
			ans.WriteByte(s[k*base+i])
			if ((k+1)*base-i)%(numRows-1) > 0 && (k+1)*base-i < n {
				ans.WriteByte(s[(k+1)*base-i])
			}
			k++
		}
	}
	return ans.String()
}