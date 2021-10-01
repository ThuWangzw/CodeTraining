// 要先深入理解进制转换的本质，就可以理解为什么每次要减一了
func convertToTitle(columnNumber int) string {
	res := ""
	for columnNumber > 0 {
		res = string(byte((columnNumber-1)%26)+'A') + res
		columnNumber = (columnNumber - 1) / 26
	}
	return res
}