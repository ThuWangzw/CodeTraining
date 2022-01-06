func simplifyPath(path string) string {
	files := strings.Split(path, "/")
	stack := make([]string, 0)
	for _, file := range files {
		if file == "" || file == "." {
			continue
		} else if file == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, file)
		}
	}
	return "/" + strings.Join(stack, "/")
}