func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	} else if num%2 == 1 {
		return 1 + numberOfSteps(num-1)
	} else {
		return 1 + numberOfSteps(num/2)
	}
}