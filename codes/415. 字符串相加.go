func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func addStrings(num1 string, num2 string) string {
	ai := len(num1)
	bi := len(num2)
	i := 0
	res := ""
	c := 0
	for i < max(ai, bi) {
		a := 0
		if ai-1-i >= 0 {
			a = int(num1[ai-1-i] - '0')
		}
		b := 0
		if bi-i-1 >= 0 {
			b = int(num2[bi-1-i] - '0')
		}

		res = strconv.Itoa((a+b+c)%10) + res
		c = (a + b + c) / 10
		i++
	}
	if c > 0 {
		res = strconv.Itoa(c) + res
	}
	return res
}