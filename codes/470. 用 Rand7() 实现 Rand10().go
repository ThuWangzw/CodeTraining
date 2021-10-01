// 拒绝采样法
func rand10() int {
	res := 40
	for res > 39 {
		res = rand7()*7 - rand7()
	}
	return res/4 + 1
}