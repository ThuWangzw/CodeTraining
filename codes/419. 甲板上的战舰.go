func countBattleships(board [][]byte) int {
	cnt := 0
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				cnt++
				if j < n-1 && board[i][j+1] == 'X' {
					for k := j; k < n && board[i][k] == 'X'; k++ {
						board[i][k] = '.'
					}
				} else {
					for k := i; k < m && board[k][j] == 'X'; k++ {
						board[k][j] = '.'
					}
				}
			}
		}
	}
	return cnt
}