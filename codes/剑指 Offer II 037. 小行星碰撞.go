func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for _, num := range asteroids {
		if len(stack) == 0 {
			stack = append(stack, num)
			continue
		}
		in := true
		for len(stack) > 0 {
			n := len(stack)
			top := stack[n-1]
			if top*num > 0 || top < 0 {
				in = true
				break
			} else if abs(top) < abs(num) {
				stack = stack[:n-1]
				in = true
			} else if abs(top) > abs(num) {
				in = false
				break
			} else {
				stack = stack[:n-1]
				in = false
				break
			}
		}
		if in {
			stack = append(stack, num)
		}
	}
	return stack
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}