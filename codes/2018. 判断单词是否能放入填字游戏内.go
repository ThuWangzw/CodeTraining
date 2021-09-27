func search(board [][]byte, status [][]int8, word string, i int, j int, direction int) bool {
	m := len(board)
	n := len(board[0])
	wordlen := len(word)
	if direction == 1 {
		for k := 0; k < wordlen; k++ {
			if i+k >= m {
				return false
			}
			status[i+k][j] = status[i+k][j] & 1
			if board[i+k][j] != ' ' && board[i+k][j] != word[k] {
				return false
			}
		}
		if i+wordlen == m || board[i+wordlen][j] == '#' {
			return true
		}
	} else if direction == 2 {
		for k := 0; k < wordlen; k++ {
			if i-k < 0 {
				return false
			}
			status[i-k][j] = status[i-k][j] & 2
			if board[i-k][j] != ' ' && board[i-k][j] != word[k] {
				return false
			}
		}
		if i-wordlen < 0 || board[i-wordlen][j] == '#' {
			return true
		}
	} else if direction == 4 {
		for k := 0; k < wordlen; k++ {
			if j+k >= n {
				return false
			}
			status[i][j+k] = status[i][j+k] & 4
			if board[i][j+k] != ' ' && board[i][j+k] != word[k] {
				return false
			}
		}
		if j+wordlen == n || board[i][j+wordlen] == '#' {
			return true
		}
	} else if direction == 8 {
		for k := 0; k < wordlen; k++ {
			if j-k < 0 {
				return false
			}
			status[i][j-k] = status[i][j-k] & 8
			if board[i][j-k] != ' ' && board[i][j-k] != word[k] {
				return false
			}
		}
		if j-wordlen < 0 || board[i][j-wordlen] == '#' {
			return true
		}
	}
	return false
}

func placeWordInCrossword(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	status := make([][]int8, m)
	for i := 0; i < m; i++ {
		status[i] = make([]int8, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' || status[i][j] == 15 {
				continue
			}
			if (status[i][j] & 1) == 0 {
				if i == 0 || board[i-1][j] == '#' {
					ans := search(board, status, word, i, j, 1)
					if ans {
						return true
					}
				}
			}
			if (status[i][j] & 4) == 0 {
				if j != 0 && board[i][j-1] != '#' {
					continue
				}
				ans := search(board, status, word, i, j, 4)
				if ans {
					return true
				}
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if board[i][j] == '#' || status[i][j] == 15 {
				continue
			}
			if (status[i][j] & 2) == 0 {
				if i == m-1 || board[i+1][j] == '#' {
					ans := search(board, status, word, i, j, 2)
					if ans {
						return true
					}
				}
			}
			if (status[i][j] & 8) == 0 {
				if j != n-1 && board[i][j+1] != '#' {
					continue
				}
				ans := search(board, status, word, i, j, 8)
				if ans {
					return true
				}
			}
		}
	}
	return false
}