func pushDominoes(dominoes string) string {
	n := len(dominoes)
	leftminright := make([]int, n)
	rightminleft := make([]int, n)
	leftclose := math.MinInt32
	for i, ch := range dominoes {
		if ch == 'L' {
			leftclose = math.MinInt32
		} else if ch == 'R' {
			leftclose = i
		} else {
			leftminright[i] = leftclose
		}
	}
	rightclose := math.MaxInt32
	for i := n - 1; i >= 0; i-- {
		ch := dominoes[i]
		if ch == 'L' {
			rightclose = i
		} else if ch == 'R' {
			rightclose = math.MaxInt32
		} else {
			rightminleft[i] = rightclose
		}
	}
	ans := make([]rune, n)
	for i, ch := range dominoes {
		if ch == 'L' || ch == 'R' {
			ans[i] = ch
		} else {
			if leftminright[i] == math.MinInt32 && rightminleft[i] == math.MaxInt32 {
				ans[i] = '.'
			} else if leftminright[i] == math.MinInt32 {
				ans[i] = 'L'
			} else if rightminleft[i] == math.MaxInt32 {
				ans[i] = 'R'
			} else {
				if i-leftminright[i] < rightminleft[i]-i {
					ans[i] = 'R'
				} else if i-leftminright[i] > rightminleft[i]-i {
					ans[i] = 'L'
				} else {
					ans[i] = '.'
				}
			}
		}
	}
	return string(ans)
}