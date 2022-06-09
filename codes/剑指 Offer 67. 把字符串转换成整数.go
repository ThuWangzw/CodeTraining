func strToInt(str string) int {
	flag := 1
	overlfow := false
	ans := 0
	n := len(str)
	// remove blank
	i := 0
	for i < n && str[i] == ' ' {
		i++
	}
	if i == n {
		return 0
	} else if str[i] == '+' {
		flag = 1
		i++
	} else if str[i] == '-' {
		flag = -1
		i++
	} else if str[i] < '0' || str[i] > '9' {
		return 0
	}
	// parse int
	for ; i < n; i++ {
		if str[i] >= '0' && str[i] <= '9' {
			ans = ans*10 + int(str[i]-'0')
			if ans > math.MaxInt32 {
				// overflow
				overlfow = true
			}
		} else {
			break
		}
	}
	// result
	if overlfow {
		if flag == 1 {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}
	return int(flag * ans)
}