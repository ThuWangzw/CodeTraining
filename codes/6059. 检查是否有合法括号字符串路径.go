func hasValidPath(grid [][]byte) bool {
	m := len(grid)
	n := len(grid[0])
	if (m+n)%2 == 0 || grid[0][0] != '(' {
		return false
	}
	cnt := make([][]map[int]bool, m)
	for i := 0; i < m; i++ {
		cnt[i] = make([]map[int]bool, n)
		for j := 0; j < n; j++ {
			cnt[i][j] = make(map[int]bool)
		}
	}
	cnt[0][0][1] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j != 0 {
				for k, _ := range cnt[i][j-1] {
					if grid[i][j] == '(' {
						cnt[i][j][k+1] = true
					} else if grid[i][j] == ')' && k > 0 {
						cnt[i][j][k-1] = true
					}
				}
			}
			if i != 0 {
				for k, _ := range cnt[i-1][j] {
					if grid[i][j] == '(' {
						cnt[i][j][k+1] = true
					} else if grid[i][j] == ')' && k > 0 {
						cnt[i][j][k-1] = true
					}
				}
			}
		}
	}
	return cnt[m-1][n-1][0]
}