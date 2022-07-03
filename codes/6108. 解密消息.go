func decodeMessage(key string, message string) string {
	cmap := make(map[byte]byte)
	for i := 0; i < len(key); i++ {
		if key[i] >= 'a' && key[i] <= 'z' {
			if _, ok := cmap[key[i]]; !ok {
				cmap[key[i]] = 'a' + byte(len(cmap))
			}
		}
	}
	newMsg := make([]byte, 0, len(message))
	for i := 0; i < len(message); i++ {
		if c, ok := cmap[message[i]]; ok {
			newMsg = append(newMsg, c)
		} else {
			newMsg = append(newMsg, message[i])
		}
	}
	return string(newMsg)
}