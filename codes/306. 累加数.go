//大整数加法，只要枚举序列的前两个数就可以判断是否是累加数序列
func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func addStr(a string, b string) string {
	m := len(a)
	n := len(b)
	last := 0
	c := make([]byte, 0)
	for i := 0; i < max(m, n); i++ {
		var ai int
		if m-1-i >= 0 {
			ai = int(a[m-1-i] - '0')
		}
		var bi int
		if n-1-i >= 0 {
			bi = int(b[n-1-i] - '0')
		}
		res := ai + bi + last
		c = append(c, byte(res%10+'0'))
		last = res / 10
	}
	if last > 0 {
		c = append(c, byte(last+'0'))
	}
	for i := 0; i < len(c)/2; i++ {
		c[i], c[len(c)-1-i] = c[len(c)-1-i], c[i]
	}
	return string(c)
}

func isAdditiveNumber(num string) bool {
	n := len(num)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a := num[:i+1]
			b := num[i+1 : j+1]
			ci := j + 1
			for {
				if (len(a) > 1 && a[0] == '0') || (len(b) > 1 && b[0] == '0') {
					break
				}
				c := addStr(a, b)
				m := len(c)
				if ci+m <= n && c == num[ci:ci+m] {
					a = b
					b = c
					ci = ci + m
					if ci == n {
						return true
					}
				} else {
					break
				}
			}
		}
	}
	return false
}