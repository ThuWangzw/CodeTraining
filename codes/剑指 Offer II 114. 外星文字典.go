func alienOrder(words []string) string {
	graph := make(map[byte]map[byte]bool)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			if _, ok := graph[word[i]]; !ok {
				graph[word[i]] = make(map[byte]bool)
			}
		}
	}
	for i := 1; i < len(words); i++ {
		worda, wordb := words[i-1], words[i]
		j := 0
		m, n := len(worda), len(wordb)
		for j < m && j < n {
			if worda[j] == wordb[j] {
				j++
			} else {
				ca, cb := worda[j], wordb[j]
				graph[ca][cb] = true
				if graph[cb][ca] {
					return ""
				}
				break
			}
		}
		if j == m || j == n {
			if m > n {
				return ""
			}
		}
	}
	queue := make([]byte, 0)
	parentCnt := make(map[byte]int)
	ans := strings.Builder{}
	for _, children := range graph {
		for child, _ := range children {
			parentCnt[child]++
		}
	}
	for c, _ := range graph {
		if parentCnt[c] == 0 {
			queue = append(queue, c)
		}
	}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		ans.WriteByte(top)
		for child, _ := range graph[top] {
			parentCnt[child]--
			if parentCnt[child] == 0 {
				queue = append(queue, child)
			}
		}
	}
	if ans.Len() != len(graph) {
		return ""
	}
	return ans.String()
}