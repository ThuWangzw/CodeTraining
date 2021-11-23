func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	scnt := make(map[byte]int)
	flag := false
	diffCnt := 0
	diffPos := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if !flag {
			scnt[s[i]]++
			if scnt[s[i]] > 1 {
				flag = true
			}
		}

		if s[i] != goal[i] {
			diffCnt++
			diffPos = append(diffPos, i)
		}
	}
	if (diffCnt == 2 && s[diffPos[0]] == goal[diffPos[1]] && s[diffPos[1]] == goal[diffPos[0]]) || (diffCnt == 0 && flag) {
		return true
	}
	return false
}