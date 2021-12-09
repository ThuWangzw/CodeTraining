func isWin(board []string) (bool, bool) {
	xwin := false
	owin := false
	for row := 0; row < 3; row++ {
		if board[row][0] == 'X' && board[row][1] == 'X' && board[row][2] == 'X' {
			xwin = true
		}
		if board[row][0] == 'O' && board[row][1] == 'O' && board[row][2] == 'O' {
			owin = true
		}
	}
	for col := 0; col < 3; col++ {
		if board[0][col] == 'X' && board[1][col] == 'X' && board[2][col] == 'X' {
			xwin = true
		}
		if board[0][col] == 'O' && board[1][col] == 'O' && board[2][col] == 'O' {
			owin = true
		}
	}
	if board[0][0] == 'X' && board[1][1] == 'X' && board[2][2] == 'X' {
		xwin = true
	}
	if board[0][2] == 'X' && board[1][1] == 'X' && board[2][0] == 'X' {
		xwin = true
	}
	if board[0][0] == 'O' && board[1][1] == 'O' && board[2][2] == 'O' {
		owin = true
	}
	if board[0][2] == 'O' && board[1][1] == 'O' && board[2][0] == 'O' {
		owin = true
	}
	return xwin, owin
}

func validTicTacToe(board []string) bool {
	xcount := 0
	ocount := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 'X' {
				xcount++
			} else if board[i][j] == 'O' {
				ocount++
			}
		}
	}
	if xcount != ocount && xcount != ocount+1 {
		return false
	}
	xwin, owin := isWin(board)
	if !xwin && !owin {
		return true
	}
	if xwin && !owin && xcount == ocount+1 {
		return true
	}
	if !xwin && owin && xcount == ocount {
		return true
	}
	return false
}