func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func computeScore(children [][]int, scores []int, target int) int {
	score := 1
	for _, child := range children[target] {
		score += computeScore(children, scores, child)
	}
	scores[target] = score
	return score
}

func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	children := make([][]int, n)
	for i := 0; i < n; i++ {
		children[i] = make([]int, 0)
	}
	for i, parent := range parents {
		if parent != -1 {
			children[parent] = append(children[parent], i)
		}
	}

	nodeweights := make([]int, n)
	computeScore(children, nodeweights, 0)
	scores := make([]int, n)
	maxscore := 0
	for i := 0; i < n; i++ {
		score := 1
		for _, child := range children[i] {
			score *= nodeweights[child]
		}
		if n-nodeweights[i] > 0 {
			score *= n - nodeweights[i]
		}
		maxscore = max(maxscore, score)
		scores[i] = score
	}
	cnt := 0
	for _, score := range scores {
		if score == maxscore {
			cnt++
		}
	}
	return cnt
}