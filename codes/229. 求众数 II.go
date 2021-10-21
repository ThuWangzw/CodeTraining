// 摩尔投票法的变形
func majorityElement(nums []int) []int {
	a := 0
	ac := 0
	b := 0
	bc := 0
	for _, num := range nums {
		if a == num {
			ac++
		} else if b == num {
			bc++
		} else if ac == 0 {
			a = num
			ac = 1
		} else if bc == 0 {
			b = num
			bc = 1
		} else {
			ac--
			bc--
		}
	}
	ans := []int{}
	acc := 0
	bcc := 0
	for _, num := range nums {
		if num == a {
			acc++
		}
		if num == b {
			bcc++
		}
	}
	n := len(nums)
	if ac != 0 && acc > n/3 {
		ans = append(ans, a)
	}
	if bc != 0 && bcc > n/3 {
		ans = append(ans, b)
	}
	return ans
}