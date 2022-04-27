func replaceSpace(s string) string {
	builder := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			builder.WriteByte(s[i])
		} else {
			builder.WriteString("%20")
		}
	}
	return builder.String()
}