var ans []string

func restoreIpIter(s string, half []byte, i int) {
	if i == 0 {
		if len(s) > 0 {
			return
		}
		ans = append(ans, string(half))
		return
	}
	if len(s) == 0 {
		return
	}
	tmp := half
	if len(half) > 0 {
		tmp = append(half, '.')
	}
	num := 0
	for j := 0; j < 3 && j < len(s); j++ {
		if j == 0 && s[j] == '0' {
			restoreIpIter(s[1:], append(tmp, '0'), i-1)
			break
		}
		num = num*10 + int(s[j]-'0')
		if num <= 255 {
			tmp = append(tmp, s[j])
			restoreIpIter(s[j+1:], tmp, i-1)
		}
	}
}

func restoreIpAddresses(s string) []string {
	ans = make([]string, 0)
	restoreIpIter(s, []byte{}, 4)
	return ans
}