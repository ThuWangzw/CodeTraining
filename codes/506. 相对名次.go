func qsortIndex(score []int, index []int) {
	p := score[index[0]]
	n := len(score) - 1
	left := 1
	right := n - 1
	for left <= right {
		if score[index[left]] <= p {
			left++
		} else {
			score[index[left]], score[index[right]] = score[index[right]], score[index[left]]
			right--
		}
	}
	score[index[right]], score[index[0]] = score[index[0]], score[index[right]]
	qsortIndex(score, index[:right])
	qsortIndex(score, index[left:])
}

func findRelativeRanks(score []int) []string {
	n := len(score)
	index := make([]int, n)
	for i := 0; i < n; i++ {
		index[i] = i
	}
	qsortIndex(score, index)
	ans := make([]string, n)
	for i, idx := range index {
		if i == 0 {
			ans[idx] = "Gold Medal"
		} else if i == 1 {
			ans[idx] = "Silver Medal"
		} else if i == 2 {
			ans[idx] = "Bronze Medal"
		} else {
			ans[idx] = strconv.Itoa(i)
		}
	}
	return ans
}