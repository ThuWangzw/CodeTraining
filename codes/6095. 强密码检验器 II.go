func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	symbols := []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '+'}
	symbolMap := make(map[rune]bool)
	for _, c := range symbols {
		symbolMap[c] = true
	}
	has := make([]bool, 4)
	for i, c := range password {
		if c >= 'a' && c <= 'z' {
			has[0] = true
		}
		if c >= 'A' && c <= 'Z' {
			has[1] = true
		}
		if c >= '0' && c <= '9' {
			has[2] = true
		}
		if symbolMap[c] {
			has[3] = true
		}
		if i > 0 && password[i] == password[i-1] {
			return false
		}
	}
	return has[0] && has[1] && has[2] && has[3]
}