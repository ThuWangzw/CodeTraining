// 单调栈，从右向左
type temp struct {
	temperature int
	index       int
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ms := make([]temp, 0)
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(ms) > 0 && ms[len(ms)-1].temperature <= temperatures[i] {
			ms = ms[:len(ms)-1]
		}
		if len(ms) == 0 {
			ans[i] = 0
		} else {
			ans[i] = ms[len(ms)-1].index - i
		}
		ms = append(ms, temp{temperature: temperatures[i], index: i})
	}
	return ans
}