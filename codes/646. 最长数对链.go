// 与最长递增子序列是一种解法，贪心+二分
func sortPairs(pairs [][]int) {
	n := len(pairs)
	if n < 2 {
		return
	} else {
		pair := pairs[0]
		i := 1
		j := n - 1
		for i <= j {
			if pairs[i][0] <= pair[0] {
				i += 1
			} else {
				pairs[i], pairs[j] = pairs[j], pairs[i]
				j -= 1
			}
		}
		pairs[0], pairs[i-1] = pairs[i-1], pairs[0]
		sortPairs(pairs[:i-1])
		sortPairs(pairs[i:])
	}
}

func findLongestChain(pairs [][]int) int {
	n := len(pairs)
	sortPairs(pairs)
	longestChains := make([][]int, 0)
	longestChains = append(longestChains, pairs[0])
	for i := 1; i < n; i++ {
		cur := pairs[i]
		if cur[0] > longestChains[len(longestChains)-1][1] {
			longestChains = append(longestChains, cur)
		} else {
			i := 0
			j := len(longestChains) - 1
			for i < j {
				mid := (i + j) / 2
				if longestChains[mid][1] >= cur[0] {
					j = mid
				} else {
					i = mid + 1
				}
			}
			if (i == 0 && longestChains[i][1] > cur[1]) || (i != 0 && longestChains[i-1][1] < cur[0] && longestChains[i][1] > cur[1]) {
				longestChains[i] = cur
			}
		}
	}
	return len(longestChains)
}