func findWord(board [][]byte, word string, s, e int) bool {
	m := len(board)
	n := len(board[0])
	if len(word) == 0 {
		return true
	}
	if board[s][e] != word[0] {
		return false
	}
	if len(word) == 1 {
		return true
	}
	c := board[s][e]
	board[s][e] = 0
	defer func() {
		board[s][e] = c
	}()
	directions := [][]int{{s, e - 1}, {s, e + 1}, {s - 1, e}, {s + 1, e}}
	for _, direction := range directions {
		if direction[0] >= 0 && direction[0] < m && direction[1] >= 0 && direction[1] < n {
			if findWord(board, word[1:], direction[0], direction[1]) {
				return true
			}
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if findWord(board, word, i, j) {
				return true
			}
		}
	}
	return false
}