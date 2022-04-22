// 将字符串扩为两倍然后用kmp
func rotateString(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}