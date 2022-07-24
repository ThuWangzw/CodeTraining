func equalPairs(grid [][]int) int {
	rowmap := make(map[string]int)
	n := len(grid)
	for i := 0; i < n; i++ {
		rowmap[num2str(grid, true, i)]++
	}
	cnt := 0
	for i := 0; i < n; i++ {
		cnt += rowmap[num2str(grid, false, i)]
	}
	return cnt
}

func num2str(grid [][]int, row bool, idx int) string {
	builder := strings.Builder{}
	n := len(grid)
	if row {
		for i := 0; i < n; i++ {
			builder.WriteString(strconv.Itoa(grid[idx][i]))
			builder.WriteByte(',')
		}
	} else {
		for i := 0; i < n; i++ {
			builder.WriteString(strconv.Itoa(grid[i][idx]))
			builder.WriteByte(',')
		}
	}
	return builder.String()
}