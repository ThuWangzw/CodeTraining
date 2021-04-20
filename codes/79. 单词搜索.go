func search(board [][]byte, word string, traveled [][]bool, cur []int) bool {
	if len(word)==0 {
		return true
	}
	c := word[0]
	m := len(board)
	n := len(board[0])
	if cur[0]>0 && !traveled[cur[0]-1][cur[1]] && board[cur[0]-1][cur[1]]==c {
		traveled[cur[0]-1][cur[1]] = true
		if search(board, word[1:], traveled, []int{cur[0]-1, cur[1]}) {
			return true
		} else {
			traveled[cur[0]-1][cur[1]] = false
		}
	}
	if cur[0]<m-1 && !traveled[cur[0]+1][cur[1]] && board[cur[0]+1][cur[1]]==c {
		traveled[cur[0]+1][cur[1]] = true
		if search(board, word[1:], traveled, []int{cur[0]+1, cur[1]}) {
			return true
		} else {
			traveled[cur[0]+1][cur[1]] = false
		}
	}
	if cur[1]>0 && !traveled[cur[0]][cur[1]-1] && board[cur[0]][cur[1]-1]==c {
		traveled[cur[0]][cur[1]-1] = true
		if search(board, word[1:], traveled, []int{cur[0], cur[1]-1}) {
			return true
		} else {
			traveled[cur[0]][cur[1]-1] = false
		}
	}
	if cur[1]<n-1 && !traveled[cur[0]][cur[1]+1] && board[cur[0]][cur[1]+1]==c {
		traveled[cur[0]][cur[1]+1] = true
		if search(board, word[1:], traveled, []int{cur[0], cur[1]+1}) {
			return true
		} else {
			traveled[cur[0]][cur[1]+1] = false
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	traveled := make([][]bool, m)
	for i:=0; i<m; i++ {
		traveled[i] = make([]bool, n)
	}
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if board[i][j] == word[0] {
				traveled[i][j] = true
				if search(board, word[1:], traveled, []int{i, j}) {
					return true
				} else {
					traveled[i][j] = false
				}
			}
		}
	}
	return false
}