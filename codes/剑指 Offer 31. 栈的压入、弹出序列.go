func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	popj := 0
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[popj] {
			stack = stack[:len(stack)-1]
			popj++
		}
	}
	return len(stack) == 0
}