func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	field := make([][]int, n)
	for i := 0; i < n; i++ {
		field[i] = make([]int, n)
	}
	for id, artifact := range artifacts {
		for i := artifact[0]; i <= artifact[2]; i++ {
			for j := artifact[1]; j <= artifact[3]; j++ {
				field[i][j] = id + 1
			}
		}
	}
	digNum := make([]int, len(artifacts))
	ans := 0
	for _, cell := range dig {
		if field[cell[0]][cell[1]] > 0 {
			id := field[cell[0]][cell[1]] - 1
			digNum[id]++
			if digNum[id] == (artifacts[id][2]-artifacts[id][0]+1)*(artifacts[id][3]-artifacts[id][1]+1) {
				ans++
			}
		}
	}
	return ans
}