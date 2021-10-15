// 功能实现题，他说啥你干啥就行
func describe(str []byte) []byte {
	if len(str) == 0 {
		return []byte{}
	}
	cur := str[0]
	cnt := 0
	ans := make([]byte, 0)
	for _, ch := range str {
		if ch == cur {
			cnt++
		} else {
			ans = append(ans, []byte(strconv.Itoa(cnt))...)
			ans = append(ans, cur)
			cur = ch
			cnt = 1
		}
	}
	ans = append(ans, []byte(strconv.Itoa(cnt))...)
	ans = append(ans, cur)
	return ans
}

func countAndSay(n int) string {
	ans := []byte{'1'}
	for i := 2; i <= n; i++ {
		ans = describe(ans)
	}
	return string(ans)
}