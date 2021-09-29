func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	nums := ""
	flag := ""
	if num < 0 {
		flag = "-"
		num = -num
	}
	for num > 0 {
		nums = strconv.Itoa(num%7) + nums
		num /= 7
	}
	return flag + nums
}